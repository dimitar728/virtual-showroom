package routes

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/handlers"
	"github.com/dimitar728/virtual-showroom/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	a := handlers.AuthHandler{DB: db}
	s := handlers.ShowroomHandler{DB: db}
	b := handlers.BookingHandler{DB: db}

	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", a.Register)
	auth.POST("/login", a.Login)
	auth.GET("/me", middleware.JWTMiddleware(), a.Me)

	// public
	api.GET("/showrooms", s.List)
	api.GET("/showrooms/:id", s.Get)

	// admin
	admin := api.Group("/showrooms")
	admin.Use(middleware.JWTMiddleware(), middleware.RequireAdmin())
	admin.POST("", s.Create)
	admin.PATCH(":id", s.Update)
	admin.DELETE(":id", s.Delete)

	// bookings
	book := api.Group("/bookings")
	book.Use(middleware.JWTMiddleware())
	book.GET("/me", b.My)
	book.POST("", b.Create)
	book.PATCH(":id/cancel", b.Cancel)

	adm := api.Group("/admin")
	adm.Use(middleware.JWTMiddleware(), middleware.RequireAdmin())
	adm.GET("/bookings", b.AdminList)
}
