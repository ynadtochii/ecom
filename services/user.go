package services

import (
	"github.com/ynadtochii/ecom/db/models"
	"github.com/ynadtochii/ecom/db/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) (*models.User, error){
	return s.Repo.DeleteUser(id)
}
