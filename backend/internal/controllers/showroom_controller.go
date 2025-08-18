package controllers

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /api/showrooms
func GetShowrooms(c *fiber.Ctx) error {
	var showrooms []models.Showroom
	if err := database.DB.Find(&showrooms).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch showrooms"})
	}
	return c.JSON(showrooms)
}

// GET /api/showrooms/:id
func GetShowroomByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var showroom models.Showroom

	if err := database.DB.First(&showroom, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "showroom not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch showroom"})
	}
	return c.JSON(showroom)
}

// POST /api/showrooms (admin only)
func CreateShowroom(c *fiber.Ctx) error {
	var body models.Showroom
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if body.Name == "" || body.ModelPath == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name and model_path required"})
	}

	// Extract admin ID from context (set by middleware)
	adminID := c.Locals("userID").(uuid.UUID)
	body.CreatedBy = adminID

	if err := database.DB.Create(&body).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create showroom"})
	}
	return c.Status(201).JSON(body)
}

// PATCH /api/showrooms/:id (admin only)
func UpdateShowroom(c *fiber.Ctx) error {
	id := c.Params("id")
	var showroom models.Showroom

	if err := database.DB.First(&showroom, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "showroom not found"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := database.DB.Model(&showroom).Updates(body).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update showroom"})
	}
	return c.JSON(showroom)
}

// DELETE /api/showrooms/:id (admin only)
func DeleteShowroom(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Showroom{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete showroom"})
	}
	return c.SendStatus(204)
}
