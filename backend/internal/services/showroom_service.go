package services

import (
	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/repositories"
	"github.com/google/uuid"
)

type ShowroomService struct {
	repo *repositories.ShowroomRepository
}

func NewShowroomService(repo *repositories.ShowroomRepository) *ShowroomService {
	return &ShowroomService{repo: repo}
}

func (s *ShowroomService) Create(input *models.Showroom, createdBy uuid.UUID) error {
	showroom := &models.Showroom{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
		ModelPath:   input.ModelPath,
		Capacity:    input.Capacity,
		CreatedBy:   createdBy,
	}
	return s.repo.Create(showroom)
}

func (s *ShowroomService) GetAll() ([]models.Showroom, error) {
	return s.repo.GetAll()
}

func (s *ShowroomService) GetByID(id string) (*models.Showroom, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(uid)
}

func (s *ShowroomService) Update(showroom *models.Showroom) error {
	return s.repo.Update(showroom)
}

func (s *ShowroomService) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(uid)
}
