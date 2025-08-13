package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/internal/repositories"
	"github.com/dimitar728/virtual-showroom/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// helper to create an in-memory app with routes and a fake auth middleware
func setupTestApp(t *testing.T) (*gin.Engine, *gorm.DB, uuid.UUID, uuid.UUID) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	// seed a user and showroom
	userID := uuid.New()
	showroomID := uuid.New()
	_ = db.Create(&models.User{ID: userID, Email: "a@b.com", PasswordHash: "x"}).Error
	_ = db.Create(&models.Showroom{ID: showroomID, Name: "Room A", Capacity: 1}).Error

	r := gin.New()
	// fake auth: always inject currentUser into context
	r.Use(func(c *gin.Context) {
		c.Set("currentUser", models.User{ID: userID, Role: "user"})
		c.Next()
	})
	// capacity resolver
	getCap := func(id uuid.UUID) (int, error) {
		var s models.Showroom
		if err := db.First(&s, "id = ?", id).Error; err != nil {
			return 1, nil
		}
		if s.Capacity <= 0 {
			return 1, nil
		}
		return s.Capacity, nil
	}
	RegisterBookingRoutes(r, db, getCap)

	return r, db, userID, showroomID
}

func doJSON(t *testing.T, r http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("encode: %v", err)
		}
	}
	req, _ := http.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestBookAndListAndCancel(t *testing.T) {
	r, db, _, showroomID := setupTestApp(t)

	slot := time.Now().Add(24 * time.Hour).UTC().Truncate(time.Minute)

	// book
	w := doJSON(t, r, "POST", "/api/bookings", map[string]any{
		"showroom_id": showroomID.String(),
		"slot_time":   slot.Format(time.RFC3339),
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// list mine
	w = doJSON(t, r, "GET", "/api/bookings/me", nil)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got []models.Booking
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 booking, got %d", len(got))
	}

	// cancel
	id := got[0].ID
	w = doJSON(t, r, "PATCH", "/api/bookings/"+id.String()+"/cancel", nil)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// verify status
	var b models.Booking
	if err := db.First(&b, "id = ?", id).Error; err != nil {
		t.Fatalf("fetch booking: %v", err)
	}
	if b.Status != models.BookingStatusCancelled {
		t.Fatalf("expected cancelled, got %s", b.Status)
	}
}

func TestDoubleBookingPrevention(t *testing.T) {
	r, _, _, showroomID := setupTestApp(t)
	slot := time.Now().Add(48 * time.Hour).UTC().Truncate(time.Minute)

	// first booking ok
	w := doJSON(t, r, "POST", "/api/bookings", map[string]any{
		"showroom_id": showroomID.String(),
		"slot_time":   slot.Format(time.RFC3339),
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// second booking in same slot should 409 (capacity=1 in seed)
	w = doJSON(t, r, "POST", "/api/bookings", map[string]any{
		"showroom_id": showroomID.String(),
		"slot_time":   slot.Format(time.RFC3339),
	})
	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409 conflict, got %d", w.Code)
	}
}

func TestOnlyOwnerCanCancel(t *testing.T) {
	// set up
	gin.SetMode(gin.TestMode)
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	_ = db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{})
	userA := uuid.New()
	userB := uuid.New()
	showroomID := uuid.New()
	_ = db.Create(&models.User{ID: userA, Email: "a@b.com", PasswordHash: "x"}).Error
	_ = db.Create(&models.User{ID: userB, Email: "b@b.com", PasswordHash: "x"}).Error
	_ = db.Create(&models.Showroom{ID: showroomID, Name: "Room A", Capacity: 2}).Error

	repo := repositories.NewBookingRepository(db)
	svc := services.NewBookingService(repo)

	// create booking owned by A
	slot := time.Now().Add(2 * time.Hour).UTC().Truncate(time.Minute)
	b, err := svc.Book(nil, userA, showroomID, slot, 2)
	if err != nil {
		t.Fatalf("book: %v", err)
	}

	// app with userB signed in
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("currentUser", models.User{ID: userB, Role: "user"})
		c.Next()
	})
	RegisterBookingRoutes(r, db, func(uuid.UUID) (int, error) { return 2, nil })

	// try to cancel as B -> 403
	w := doJSON(t, r, "PATCH", "/api/bookings/"+b.ID.String()+"/cancel", nil)
	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", w.Code)
	}
}
Wiring (snippet for cmd/main.go)
go
Copy
Edit
// after router & db init
// import ( "your/module/internal/handlers"; "your/module/internal/models"; "github.com/google/uuid" )

handlers.RegisterBookingRoutes(r, db, func(id uuid.UUID) (int, error) {
	var s models.Showroom
	if err := db.First(&s, "id = ?", id).Error; err != nil {
		return 1, nil
	}
	if s.Capacity <= 0 {
		return 1, nil
	}
	return s.Capacity, nil
})