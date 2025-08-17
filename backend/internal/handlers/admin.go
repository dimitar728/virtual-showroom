package handlers

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Suspend a user
func SuspendUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("status", models.StatusSuspended).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to suspend user"})
	}
	return c.JSON(fiber.Map{"message": "User suspended"})
}

// Reactivate a user
func ReactivateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("status", models.StatusActive).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to reactivate user"})
	}
	return c.JSON(fiber.Map{"message": "User reactivated"})
}

// Delete a user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}
