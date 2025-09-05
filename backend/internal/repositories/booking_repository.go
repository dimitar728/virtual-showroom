package repositories

import (
	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *BookingRepository) GetByUser(userID uuid.UUID) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) Cancel(id uuid.UUID) error {
	return r.db.Model(&models.Booking{}).Where("id = ?", id).Update("status", "cancelled").Error
}

func (r *BookingRepository) GetAll() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) GetByUserID(userID uuid.UUID) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
