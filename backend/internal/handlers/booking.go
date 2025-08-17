package handlers

import (
	"net/http"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func isSlotAvailable(db *gorm.DB, showroomID uuid.UUID, slotTime time.Time) (bool, error) {
	var count int64
	err := db.Model(&models.Booking{}).
		Where("showroom_id = ? AND slot_time = ? AND status != ?", showroomID, slotTime, models.StatusCancelled).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func CreateBooking(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		showroomID := uuid.MustParse(req.ShowroomID)
		available, err := isSlotAvailable(db, showroomID, req.SlotTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check availability"})
			return
		}
		if !available {
			c.JSON(http.StatusConflict, gin.H{"error": "Selected slot is not available"})
			return
		}

}

func CheckAvailability(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        showroomID := c.Param("id")
        date := c.Query("date")

        day, err := time.Parse("2006-01-02", date)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date"})
            return
        }

        var bookings []models.Booking
        db.Where("showroom_id = ? AND DATE(slot_time) = ? AND status != ?", showroomID, day, models.StatusCancelled).
            Find(&bookings)

        c.JSON(http.StatusOK, gin.H{
            "showroom_id": showroomID,
            "date":        date,
            "bookedSlots": bookings,
        })
    }
}