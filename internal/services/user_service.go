package services

import (
	"errors"
	"rental-mobil/internal/models"
	"rental-mobil/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.GetUserByEmail(email)
}

func (s *UserService) UpdateUser(id uint, user *models.User) error {
	if user.Username != "" {
		existingUser, _ := s.Repo.GetUserByUsername(user.Username)
		if existingUser != nil && existingUser.ID != id {
			return errors.New("username sudah digunakan")
		}
	}
	if user.Email != "" {
		existingUser, _ := s.Repo.GetUserByEmail(user.Email)
		if existingUser != nil && existingUser.ID != id {
			return errors.New("email sudah digunakan")
		}
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email, dan password wajib diisi")
	}

	existingUser, _ := s.Repo.GetUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username sudah digunakan")
	}

	existingUser, _ = s.Repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email sudah digunakan")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.Repo.CreateUser(user)
}
