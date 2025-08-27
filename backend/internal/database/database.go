package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	URL       string
	UseSQLite bool // for tests
}

func Connect(cfg Config) *gorm.DB {
	var db *gorm.DB
	var err error
	if cfg.UseSQLite {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		url := cfg.URL
		if url == "" {
			url = os.Getenv("DATABASE_URL")
		}
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// Connection pool (for Postgres driver)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return db
}
