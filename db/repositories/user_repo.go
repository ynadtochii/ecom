package repositories

import (
	"github.com/ynadtochii/ecom/db/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
  if err := r.DB.Create(user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
  if err := r.DB.Save(user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

func (r *UserRepository) DeleteUser(id uint) (*models.User, error) {
  var user models.User
  if err := r.DB.First(&user, id).Error; err != nil {
    return nil, err
  }
  if err := r.DB.Delete(&user).Error; err != nil {
    return nil, err
  }
  return &user, nil
}
