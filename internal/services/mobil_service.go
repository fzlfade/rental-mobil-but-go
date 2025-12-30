package services

import (
	"rental-mobil/internal/models"
	"rental-mobil/internal/repositories"
)

type MobilService struct {
	Repo *repositories.MobilRepository
}

func NewMobilService(repo *repositories.MobilRepository) *MobilService {
	return &MobilService{Repo: repo}
}

func (s *MobilService) GetAllMobil() ([]models.Mobil, error) {
	return s.Repo.GetAllMobil()
}

func (s *MobilService) GetMobilByID(id uint) (*models.Mobil, error) {
	return s.Repo.GetMobilByID(id)
}

func (s *MobilService) CreateMobil(mobil *models.Mobil) error {
	return s.Repo.CreateMobil(mobil)
}

func (s *MobilService) UpdateMobil(id uint, mobil *models.Mobil) error {
	return s.Repo.UpdateMobil(id, mobil)
}

func (s *MobilService) DeleteMobil(id uint) error {
	return s.Repo.DeleteMobil(id)
}
