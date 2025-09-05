package services

import (
	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.GetUser(uid)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.userRepo.DeleteUser(uid)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.ListUsers()
}

func (s *UserService) SuspendUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.userRepo.SuspendUser(uid)
}

func (s *UserService) ReactivateUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.userRepo.ReactivateUser(uid)
}
