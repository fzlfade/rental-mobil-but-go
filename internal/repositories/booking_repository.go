package repositories

import (
	"errors"
	"rental-mobil/internal/models"

	"gorm.io/gorm"
)

type BookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (r *BookingRepository) GetBookingByID(id uint) (*models.Booking, error) {
	var b models.Booking
	err := r.DB.Preload("User").Preload("Mobil").First(&b, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &b, err
}

func (r *BookingRepository) GetBookingsByUserID(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.Preload("Mobil").Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) GetBookingsByMobilID(mobilID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.Preload("User").Where("mobil_id = ?", mobilID).Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
	return r.DB.Create(booking).Error
}

func (r *BookingRepository) UpdateBooking(id uint, booking *models.Booking) error {
	return r.DB.Model(&models.Booking{}).Where("id = ?", id).Updates(booking).Error
}

func (r *BookingRepository) DeleteBooking(id uint) error {
	return r.DB.Delete(&models.Booking{}, id).Error
}
