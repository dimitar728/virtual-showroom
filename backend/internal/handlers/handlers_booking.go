package handlers

import (
	"net/http"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/internal/repositories"
	"github.com/dimitar728/virtual-showroom/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookSlotRequest struct {
	ShowroomID string `json:"showroom_id" binding:"required,uuid"`
	SlotTime   string `json:"slot_time" binding:"required"` // ISO8601
}

// getCapacity returns capacity for a showroomID (plug your DB lookup)
func RegisterBookingRoutes(r *gin.Engine, db *gorm.DB, getCapacity func(showroomID uuid.UUID) (int, error)) {
	repo := repositories.NewBookingRepository(db)
	svc := services.NewBookingService(repo)

	// current user extractor
	currentUser := func(c *gin.Context) (uuid.UUID, bool, bool) {
		// Prefer a full user struct in context
		if v, ok := c.Get("currentUser"); ok {
			if u, ok2 := v.(models.User); ok2 {
				return u.ID, u.Role == "admin", true
			}
		}
		// Fallback to IDs set by other middleware
		if v, ok := c.Get("currentUserID"); ok {
			if id, ok2 := v.(uuid.UUID); ok2 {
				return id, false, true
			}
			if s, ok2 := v.(string); ok2 {
				if parsed, err := uuid.Parse(s); err == nil {
					return parsed, false, true
				}
			}
		}
		return uuid.Nil, false, false
	}

	// POST /api/bookings
	r.POST("/api/bookings", func(c *gin.Context) {
		uid, _, ok := currentUser(c)
		if !ok || uid == uuid.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		var req BookSlotRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		showroomID, _ := uuid.Parse(req.ShowroomID)
		slot, err := time.Parse(time.RFC3339, req.SlotTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "slot_time must be RFC3339"})
			return
		}

		capacity := 1
		if getCapacity != nil {
			if cap, err := getCapacity(showroomID); err == nil && cap > 0 {
				capacity = cap
			}
		}

		booking, err := svc.Book(c.Request.Context(), uid, showroomID, slot, capacity)
		if err != nil {
			switch err {
			case repositories.ErrSlotFull:
				c.JSON(http.StatusConflict, gin.H{"error": "slot is full"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusCreated, booking)
	})

	// GET /api/bookings/me
	r.GET("/api/bookings/me", func(c *gin.Context) {
		uid, _, ok := currentUser(c)
		if !ok || uid == uuid.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		list, err := svc.ListMine(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// PATCH /api/bookings/:id/cancel
	r.PATCH("/api/bookings/:id/cancel", func(c *gin.Context) {
		uid, isAdmin, ok := currentUser(c)
		if !ok || uid == uuid.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		if err := svc.Cancel(c.Request.Context(), id, uid, isAdmin); err != nil {
			switch err {
			case repositories.ErrBookingNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			case repositories.ErrUnauthorizedOwner:
				c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}
		c.Status(http.StatusNoContent)
	})
}
