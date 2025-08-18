package handlers

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // Optional: Only admins can set
}

func Register(c *fiber.Ctx) error {
	var body RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		Email:        body.Email,
		PasswordHash: string(hash),
		Role:         models.RoleUser,
	}

	// Optional: Allow role assignment only for admin
	if body.Role == "admin" {
		// Check if requester is admin
		reqUser := c.Locals("user")
		if reqUser != nil && reqUser.(string) == "admin" {
			user.Role = models.RoleAdmin
		}
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}
