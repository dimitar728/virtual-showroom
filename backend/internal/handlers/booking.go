package handlers

import (
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingHandler struct{ DB *gorm.DB }

type bookReq struct {
	ShowroomID uuid.UUID `json:"showroom_id" binding:"required"`
	SlotTime   time.Time `json:"slot_time" binding:"required"`
}

func (h BookingHandler) My(c *gin.Context) {
	uid := uuid.MustParse(c.GetString("uid"))
	var b []models.Booking
	h.DB.Where("user_id = ?", uid).Order("slot_time asc").Find(&b)
	c.JSON(200, b)
}

func (h BookingHandler) Create(c *gin.Context) {
	var req bookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// prevent double-booking: same showroom and slot
	var cnt int64
	h.DB.Model(&models.Booking{}).Where("showroom_id = ? AND slot_time = ? AND status != ?", req.ShowroomID, req.SlotTime, models.BookingCancelled).Count(&cnt)
	if cnt > 0 {
		c.JSON(409, gin.H{"error": "slot already booked"})
		return
	}
	b := models.Booking{
		UserID:     uuid.MustParse(c.GetString("uid")),
		ShowroomID: req.ShowroomID,
		SlotTime:   req.SlotTime,
		Status:     models.BookingPending,
	}
	if err := h.DB.Create(&b).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to book"})
		return
	}
	c.JSON(201, b)
}

func (h BookingHandler) Cancel(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	uid := uuid.MustParse(c.GetString("uid"))
	var b models.Booking
	if err := h.DB.First(&b, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	// allow owner or admin to cancel
	role := c.GetString("role")
	if role != "admin" && b.UserID != uid {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}
	if err := h.DB.Model(&b).Update("status", models.BookingCancelled).Error; err != nil {
		c.JSON(500, gin.H{"error": "cancel failed"})
		return
	}
	c.JSON(200, b)
}

func (h BookingHandler) AdminList(c *gin.Context) {
	var b []models.Booking
	h.DB.Order("created_at desc").Find(&b)
	c.JSON(200, b)
}
