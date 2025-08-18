package controllers

import (
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /api/bookings/me
func GetMyBookings(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	var bookings []models.Booking
	if err := database.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch bookings"})
	}
	return c.JSON(bookings)
}

// POST /api/bookings
func CreateBooking(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var body struct {
		ShowroomID uuid.UUID `json:"showroom_id"`
		SlotTime   time.Time `json:"slot_time"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	// check if slot already booked
	var existing models.Booking
	err := database.DB.Where("showroom_id = ? AND slot_time = ? AND status != ?", body.ShowroomID, body.SlotTime, "cancelled").First(&existing).Error
	if err != gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{"error": "slot already booked"})
	}

	booking := models.Booking{
		UserID:     userID,
		ShowroomID: body.ShowroomID,
		SlotTime:   body.SlotTime,
		Status:     "confirmed",
	}
	if err := database.DB.Create(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create booking"})
	}

	return c.Status(201).JSON(booking)
}

// PATCH /api/bookings/:id/cancel
func CancelBooking(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	id := c.Params("id")

	var booking models.Booking
	if err := database.DB.First(&booking, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "booking not found"})
	}

	// only allow owner or admin
	if booking.UserID != userID && c.Locals("role") != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "not authorized"})
	}

	booking.Status = "cancelled"
	if err := database.DB.Save(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to cancel booking"})
	}

	return c.JSON(booking)
}
