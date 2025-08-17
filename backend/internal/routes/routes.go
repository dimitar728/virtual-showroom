package routes

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/handlers"
	"github.com/dimitar728/virtual-showroom/backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	admin.Patch("/users/:id/suspend", handlers.SuspendUser)
	admin.Patch("/users/:id/reactivate", handlers.ReactivateUser)
	admin.Delete("/users/:id", handlers.DeleteUser)
}
