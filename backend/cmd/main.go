package main

import (
	"log"
	"os"

	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

func main() {
	if os.Getenv("GIN_MODE") == "" {
		os.Setenv("GIN_MODE", "release")
	}
	db := database.Connect(database.Config{URL: os.Getenv("DATABASE_URL")})
	migrate(db)

	r := gin.Default()
	routes.Register(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
