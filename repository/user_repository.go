package repository

import (
    "todolist-backend/models"
    "gorm.io/gorm"
)

type UserRepository struct {
  DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
  return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
  return r.DB.Create(user).Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
  var users []models.User
  result := r.DB.Find(&users)
  return users, result.Error
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
  var user models.User
  result := r.DB.First(&user, id)
  return user, result.Error
}

func (r *UserRepository) Update(user *models.User) error {
  return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
  return r.DB.Delete(&models.User{}, id).Error
}
