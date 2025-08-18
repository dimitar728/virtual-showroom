package routes

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/controllers"
	"github.com/dimitar728/virtual-showroom/backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func ShowroomRoutes(app *fiber.App) {
	route := app.Group("/api/showrooms")

	// Public
	route.Get("/", controllers.GetShowrooms)
	route.Get("/:id", controllers.GetShowroomByID)

	// Admin only
	route.Post("/", middleware.RequireAuth, middleware.RequireAdmin, controllers.CreateShowroom)
	route.Patch("/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.UpdateShowroom)
	route.Delete("/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeleteShowroom)
}
