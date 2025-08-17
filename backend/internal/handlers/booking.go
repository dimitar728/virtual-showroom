package handlers

import (
	"net/http"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingRequest struct {
	ShowroomID string    `json:"showroom_id" binding:"required,uuid"`
	SlotTime   time.Time `json:"slot_time" binding:"required"`
}

// POST /api/bookings
func CreateBooking(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID") // extracted from JWT middleware

		var req BookingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Prevent double booking
		var count int64
		db.Model(&models.Booking{}).
			Where("showroom_id = ? AND slot_time = ? AND status != ?", req.ShowroomID, req.SlotTime, models.StatusCancelled).
			Count(&count)

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "Slot already booked"})
			return
		}

		booking := models.Booking{
			UserID:     uuid.MustParse(userID),
			ShowroomID: uuid.MustParse(req.ShowroomID),
			SlotTime:   req.SlotTime,
			Status:     models.StatusConfirmed,
		}

		if err := db.Create(&booking).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
			return
		}

		c.JSON(http.StatusCreated, booking)
	}
}

// GET /api/bookings/me
func GetMyBookings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")

		var bookings []models.Booking
		if err := db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
			return
		}

		c.JSON(http.StatusOK, bookings)
	}
}

// PATCH /api/bookings/:id/cancel
func CancelBooking(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		bookingID := c.Param("id")

		var booking models.Booking
		if err := db.First(&booking, "id = ?", bookingID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}

		// Only owner or admin can cancel
		role := c.GetString("role")
		if booking.UserID.String() != userID && role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			return
		}

		booking.Status = models.StatusCancelled
		if err := db.Save(&booking).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
			return
		}

		c.JSON(http.StatusOK, booking)
	}
}
