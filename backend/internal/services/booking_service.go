package services

import (
	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/repositories"
	"github.com/google/uuid"
)

type BookingService struct {
	repo *repositories.BookingRepository
}

func NewBookingService(repo *repositories.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) Create(booking *models.Booking) error {
	return s.repo.Create(booking)
}

func (s *BookingService) GetByUser(userID string) ([]models.Booking, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByUser(uid)
}

func (s *BookingService) Cancel(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.Cancel(uid)
}

func (s *BookingService) GetAllBookings() ([]models.Booking, error) {
	return s.repo.GetAll()
}

func (s *BookingService) GetBookingsByUserID(userID uuid.UUID) ([]models.Booking, error) {
	return s.repo.GetByUserID(userID)
}
