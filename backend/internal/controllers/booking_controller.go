package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingController struct {
	service *services.BookingService
}

func NewBookingController(service *services.BookingService) *BookingController {
	return &BookingController{service: service}
}

func (c *BookingController) Create(ctx *gin.Context) {
	// Debug: print raw user_id from context
	rawUserID, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	fmt.Printf("[DEBUG] Raw user_id from context: %v\n", rawUserID)
	userIDStr, ok := rawUserID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID in context is not a string"})
		return
	}
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	fmt.Printf("[DEBUG] Parsed userUUID: %v\n", userUUID)
	var payload struct {
		ShowroomID string `json:"showroom_id"`
		SlotTime   string `json:"slot_time"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// userUUID already parsed above, remove duplicate assignment
	showroomUUID, err := uuid.Parse(payload.ShowroomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid showroom ID"})
		return
	}
	slotTime, err := time.Parse(time.RFC3339, payload.SlotTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slot time format"})
		return
	}
	booking := models.Booking{
		UserID:     userUUID,
		ShowroomID: showroomUUID,
		SlotTime:   slotTime,
		Status:     "pending",
	}
	if err := c.service.Create(&booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, booking)
}

func (c *BookingController) GetByUser(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	bookings, err := c.service.GetByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) Cancel(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Cancel(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Booking cancelled"})
}

func (bc *BookingController) GetAll(c *gin.Context) {
	bookings, err := bc.service.GetAllBookings()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch bookings"})
		return
	}
	c.JSON(200, bookings)
}

func (bc *BookingController) GetMyBookings(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	bookings, err := bc.service.GetBookingsByUserID(userID) // <-- fix here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
