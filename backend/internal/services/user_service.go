// internal/services/user_service.go
package services

import (
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.List()
}
