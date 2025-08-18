package middleware

import (
	"yourapp/utils"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	user, err := utils.GetUserFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	c.Locals("user", user)
	return c.Next()
}

func RequireAdmin(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil || user.(string) != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admin access required"})
	}
	return c.Next()
}
