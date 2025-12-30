package repositories

import (
	"errors"
	"rental-mobil/internal/models"
	"gorm.io/gorm"
)

type MobilRepository struct {
	DB *gorm.DB
}

func NewMobilRepository(db *gorm.DB) *MobilRepository {
	return &MobilRepository{DB: db}
}

func (r *MobilRepository) GetAllMobil() ([]models.Mobil, error) {
	var mobil []models.Mobil
	err := r.DB.Find(&mobil).Error
	return mobil, err
}

func (r *MobilRepository) GetMobilByID(id uint) (*models.Mobil, error) {
	var m models.Mobil
	err := r.DB.First(&m, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &m, err
}

func (r *MobilRepository) CreateMobil(mobil *models.Mobil) error {
	return r.DB.Create(mobil).Error
}

func (r *MobilRepository) UpdateMobil(id uint, mobil *models.Mobil) error {
	return r.DB.Model(&models.Mobil{}).Where("id = ?", id).Updates(mobil).Error
}

func (r *MobilRepository) DeleteMobil(id uint) error {
	return r.DB.Delete(&models.Mobil{}, id).Error
}