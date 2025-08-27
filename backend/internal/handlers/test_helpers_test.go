//go:build test

package handlers

import (
	"os"
	"testing"

	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTest(t *testing.T) (*gin.Engine, *gorm.DB) {
	os.Setenv("JWT_SECRET", "test-secret")
	db := database.Connect(database.Config{UseSQLite: true})
	if err := db.AutoMigrate(&models.User{}, &models.Showroom{}, &models.Booking{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	r := gin.Default()
	routes.Register(r, db)
	return r, db
}
