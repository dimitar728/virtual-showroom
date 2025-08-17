package controllers

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /api/showrooms
func CreateShowroom(c *fiber.Ctx) error {
	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
}

// PATCH /api/showrooms/:id (admin only)
func UpdateShowroom(c *fiber.Ctx) error {

	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
}
