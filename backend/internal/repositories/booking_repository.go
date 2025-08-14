package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrSlotFull          = errors.New("slot is already full")
	ErrBookingNotFound   = errors.New("booking not found")
	ErrAlreadyCancelled  = errors.New("booking already cancelled")
	ErrUnauthorizedOwner = errors.New("not owner of the booking")
)

type BookingRepository interface {
	Create(ctx context.Context, b *models.Booking, capacity int) error
	GetByUser(ctx context.Context, userID uuid.UUID) ([]models.Booking, error)
	Cancel(ctx context.Context, bookingID uuid.UUID, requesterID uuid.UUID, isAdmin bool) error
	CountConfirmed(ctx context.Context, showroomID uuid.UUID, slot time.Time) (int64, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// Create inserts a booking after verifying capacity using a SERIALIZABLE transaction.
func (r *bookingRepository) Create(ctx context.Context, b *models.Booking, capacity int) error {
	serializable := &sql.TxOptions{Isolation: sql.LevelSerializable}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Lock rows for this slot to avoid races (works on Postgres; harmless on SQLite)
		var current int64
		if err := tx.
			Model(&models.Booking{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("showroom_id = ? AND slot_time = ? AND status = ?", b.ShowroomID, b.SlotTime, models.BookingStatusConfirmed).
			Count(&current).Error; err != nil {
			return err
		}

		if current >= int64(capacity) {
			return ErrSlotFull
		}

		// Confirm immediately; unique index prevents same user same slot dup
		b.Status = models.BookingStatusConfirmed
		if err := tx.Create(&b).Error; err != nil {
			return err
		}
		return nil
	}, serializable)
}

// GetByUser lists bookings for a user ordered by time desc.
func (r *bookingRepository) GetByUser(ctx context.Context, userID uuid.UUID) ([]models.Booking, error) {
	var list []models.Booking
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("slot_time DESC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *bookingRepository) CountConfirmed(ctx context.Context, showroomID uuid.UUID, slot time.Time) (int64, error) {
	var n int64
	err := r.db.WithContext(ctx).
		Model(&models.Booking{}).
		Where("showroom_id = ? AND slot_time = ? AND status = ?", showroomID, slot, models.BookingStatusConfirmed).
		Count(&n).Error
	return n, err
}

func (r *bookingRepository) Cancel(ctx context.Context, bookingID uuid.UUID, requesterID uuid.UUID, isAdmin bool) error {
	var b models.Booking
	if err := r.db.WithContext(ctx).First(&b, "id = ?", bookingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBookingNotFound
		}
		return err
	}
	if !isAdmin && b.UserID != requesterID {
		return ErrUnauthorizedOwner
	}
	if b.Status == models.BookingStatusCancelled {
		return ErrAlreadyCancelled
	}
	return r.db.WithContext(ctx).Model(&b).Update("status", models.BookingStatusCancelled).Error
}
internal/services/booking_service.go
go
Copy
Edit
package services

import (
	"context"
	"time"

	"your/module/internal/models"
	"your/module/internal/repositories"
	"github.com/google/uuid"
)

type BookingService interface {
	Book(ctx context.Context, userID, showroomID uuid.UUID, slot time.Time, capacity int) (*models.Booking, error)
	ListMine(ctx context.Context, userID uuid.UUID) ([]models.Booking, error)
	Cancel(ctx context.Context, bookingID uuid.UUID, requesterID uuid.UUID, isAdmin bool) error
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) Book(ctx context.Context, userID, showroomID uuid.UUID, slot time.Time, capacity int) (*models.Booking, error) {
	b := &models.Booking{
		UserID:     userID,
		ShowroomID: showroomID,
		SlotTime:   slot.UTC().Truncate(time.Minute),
	}
	err := s.repo.Create(ctx, b, capacity)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *bookingService) ListMine(ctx context.Context, userID uuid.UUID) ([]models.Booking, error) {
	return s.repo.GetByUser(ctx, userID)
}

func (s *bookingService) Cancel(ctx context.Context, bookingID uuid.UUID, requesterID uuid.UUID, isAdmin bool) error {
	return s.repo.Cancel(ctx, bookingID, requesterID, isAdmin)
}