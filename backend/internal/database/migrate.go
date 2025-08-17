package database

import (
	"log"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Booking{})
}
