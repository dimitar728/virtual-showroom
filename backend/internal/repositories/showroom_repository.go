package repositories

import (
	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowroomRepository struct {
	db *gorm.DB
}

func NewShowroomRepository(db *gorm.DB) *ShowroomRepository {
	return &ShowroomRepository{db: db}
}

func (r *ShowroomRepository) Create(showroom *models.Showroom) error {
	return r.db.Create(showroom).Error
}

func (r *ShowroomRepository) GetAll() ([]models.Showroom, error) {
	var showrooms []models.Showroom
	err := r.db.Find(&showrooms).Error
	return showrooms, err
}

func (r *ShowroomRepository) GetByID(id uuid.UUID) (*models.Showroom, error) {
	var showroom models.Showroom
	err := r.db.Where("id = ?", id).First(&showroom).Error
	return &showroom, err
}

func (r *ShowroomRepository) Update(showroom *models.Showroom) error {
	return r.db.Save(showroom).Error
}

func (r *ShowroomRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Showroom{}, "id = ?", id).Error
}
