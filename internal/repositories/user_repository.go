package repositories

import (
	"errors"
	"rental-mobil/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var u models.User
	err := r.DB.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &u, err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var u models.User
	err := r.DB.Where("username = ?", username).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &u, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.DB.Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &u, err
}

func (r *UserRepository) UpdateUser(id uint, user *models.User) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
