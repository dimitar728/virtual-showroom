package routes

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/controllers"
	"github.com/dimitar728/virtual-showroom/backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookingRoutes(app *fiber.App) {
	route := app.Group("/api/bookings", middleware.RequireAuth)

	route.Get("/me", controllers.GetMyBookings)
	route.Post("/", controllers.CreateBooking)
	route.Patch("/:id/cancel", controllers.CancelBooking)
}
