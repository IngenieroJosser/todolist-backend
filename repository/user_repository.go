package repository

import (
	"errors"
	"todolist-backend/models"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	if err := user.HashPassword(); err != nil {
		return err
	}
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
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, ErrUserNotFound
	}
	return user, result.Error
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, ErrUserNotFound
	}
	return user, result.Error
}

func (r *UserRepository) Update(user *models.User) error {
	// Si se actualiza la contraseña, hashearla
	if user.Password != "" {
		if err := user.HashPassword(); err != nil {
			return err
		}
	}
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}