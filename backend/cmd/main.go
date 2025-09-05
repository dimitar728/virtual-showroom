package main

import (
	"log"

	"github.com/ajonesb/user-management/backend/internal/controllers"
	"github.com/ajonesb/user-management/backend/internal/middleware"
	"github.com/ajonesb/user-management/backend/internal/repositories"
	"github.com/ajonesb/user-management/backend/internal/services"
	"github.com/ajonesb/user-management/backend/pkg/config"
	"github.com/ajonesb/user-management/backend/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	userRepo := repositories.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization", "Content-Type", "Accept")
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/api/auth/register", authController.Register)
	r.POST("/api/auth/login", authController.Login)

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware(cfg))
	{
		authorized.GET("/auth/me", authController.Me)
		authorized.GET("/admin/users", userController.ListUsers)
		authorized.GET("/admin/users/:id", userController.GetUser)
		authorized.PATCH("/admin/users/:id", userController.UpdateUser)
		authorized.PATCH("/admin/users/:id/suspend", userController.SuspendUser)
		authorized.PATCH("/admin/users/:id/reactivate", userController.ReactivateUser)
		authorized.DELETE("/admin/users/:id", userController.DeleteUser)
	}

	showroomRepo := repositories.NewShowroomRepository(database.DB)
	showroomService := services.NewShowroomService(showroomRepo)
	showroomController := controllers.NewShowroomController(showroomService)

	bookingRepo := repositories.NewBookingRepository(database.DB)
	bookingService := services.NewBookingService(bookingRepo)
	bookingController := controllers.NewBookingController(bookingService)

	// Public Showroom routes
	r.GET("/api/showrooms", showroomController.GetAll)
	r.GET("/api/showrooms/:id", showroomController.GetByID)

	// api := r.Group("/api") // removed unused variable
	// You can add routes to 'api' here if needed

	// Admin Showroom routes
	admin := authorized.Group("/admin")
	admin.POST("/showrooms", showroomController.Create)
	admin.PATCH("/showrooms/:id", showroomController.Update)
	admin.DELETE("/showrooms/:id", showroomController.Delete)

	// Booking routes
	authorized.GET("/bookings", bookingController.GetByUser)
	authorized.POST("/bookings", bookingController.Create)
	authorized.PATCH("/bookings/:id/cancel", bookingController.Cancel)
	admin.GET("/bookings", bookingController.GetAll)

	r.Run(":8080")
}
