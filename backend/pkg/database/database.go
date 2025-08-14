// pkg/database/database.go
package database

import (
	"fmt"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }

	// // Auto Migrate the schema
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	return nil, err
	// }

	// return db, nil
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Auto Migrate your schema
	if err := db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{}); err != nil {
		return nil, err
	}

	// Unique index to prevent same user from booking same slot twice
	// and speed up lookups by slot.
	_ = db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS uq_user_showroom_slot
		ON bookings (user_id, showroom_id, slot_time);
	`).Error

	return db, nil
}
