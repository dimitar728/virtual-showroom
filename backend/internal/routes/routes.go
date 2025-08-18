package routes

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/handlers"
	"github.com/dimitar728/virtual-showroom/backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public
	api.Post("/auth/register", handlers.Register)

	// Admin only
	admin := api.Group("/admin", middleware.RequireAuth, middleware.RequireAdmin)
	admin.Get("/users", handlers.GetAllUsers)
	admin.Patch("/users/:id", handlers.UpdateUserRole)
	admin.Patch("/users/:id/suspend", handlers.SuspendUser)
	admin.Patch("/users/:id/reactivate", handlers.ReactivateUser)
	admin.Delete("/users/:id", handlers.DeleteUser)
}
