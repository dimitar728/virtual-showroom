package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ajonesb/user-management/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		// Check for missing variables
		if host == "" || port == "" || user == "" || password == "" || dbname == "" {
			log.Fatal("Database environment variables are not set correctly. Please set DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, and DB_NAME.")
		}
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db

	// Auto-migrate all models
	err = db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
