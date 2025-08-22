package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func makeUserToken(uid string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid, "role": "user", "exp": time.Now().Add(time.Hour).Unix()})
	s, _ := t.SignedString([]byte("test-secret"))
	return s
}

func seedShowroom(db *gorm.DB) (string, error) {
	s := models.Showroom{Name: "Seed", Capacity: 5}
	if err := db.Create(&s).Error; err != nil {
		return "", err
	}
	return s.ID.String(), nil
}

func TestBookingFlow(t *testing.T) {
	r, db := setupTest(t)
	roomID, err := seedShowroom(db)
	if err != nil {
		t.Fatal(err)
	}

	userToken := makeUserToken("00000000-0000-0000-0000-000000000002")

	// book
	payload := map[string]any{"showroom_id": roomID, "slot_time": time.Now().Add(2 * time.Hour)}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/bookings", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userToken)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("book expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// my bookings
	req = httptest.NewRequest(http.MethodGet, "/api/bookings/me", nil)
	req.Header.Set("Authorization", "Bearer "+userToken)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("my expected 200, got %d", w.Code)
	}
}
