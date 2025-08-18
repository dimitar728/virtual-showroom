package database

import (
	"log"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
)

func Migrate() {

	err := DB.AutoMigrate(&models.User{})
	err := DB.AutoMigrate(&models.Booking{})


	err := DB.AutoMigrate(&models.User{})
	err := DB.AutoMigrate(&models.Booking{})

	err := DB.AutoMigrate(&models.Booking{})
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("Database migrated successfully")


	return db.AutoMigrate(
		&models.User{},
		&models.Showroom{},
		&models.Booking{},
		&models.Hotspot{}, // <-- add this
	)

}
