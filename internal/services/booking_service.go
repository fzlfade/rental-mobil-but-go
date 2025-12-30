package services

import (
	"errors"
	"rental-mobil/internal/models"
	"rental-mobil/internal/repositories"

	"gorm.io/gorm"
)

type BookingService struct {
	Repo      *repositories.BookingRepository
	UserRepo  *repositories.UserRepository
	MobilRepo *repositories.MobilRepository
}

func NewBookingService(repo *repositories.BookingRepository, userRepo *repositories.UserRepository, mobilRepo *repositories.MobilRepository) *BookingService {
	return &BookingService{Repo: repo, UserRepo: userRepo, MobilRepo: mobilRepo}
}

func (s *BookingService) GetBookingByID(id uint) (*models.Booking, error) {
	return s.Repo.GetBookingByID(id)
}

func (s *BookingService) GetBookingsByUserID(userID uint) ([]models.Booking, error) {
	_, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetBookingsByUserID(userID)
}

func (s *BookingService) GetBookingsByMobilID(mobilID uint) ([]models.Booking, error) {
	_, err := s.MobilRepo.GetMobilByID(mobilID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetBookingsByMobilID(mobilID)
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	if booking.UserID == 0 || booking.MobilID == 0 {
		return errors.New("user_id dan mobil_id wajib diisi")
	}
	if booking.TanggalMulai.After(booking.TanggalSelesai) {
		return errors.New("tanggal mulai tidak boleh setelah tanggal selesai")
	}

	_, err := s.UserRepo.GetUserByID(booking.UserID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	_, err = s.MobilRepo.GetMobilByID(booking.MobilID)
	if err != nil {
		return errors.New("mobil tidak ditemukan")
	}

	var existingBooking models.Booking
	err = s.Repo.DB.Where("mobil_id = ? AND ((tanggal_mulai <= ? AND tanggal_selesai >= ?) OR (tanggal_mulai <= ? AND tanggal_selesai >= ?))",
		booking.MobilID,
		booking.TanggalSelesai, booking.TanggalMulai,
		booking.TanggalSelesai, booking.TanggalMulai,
	).First(&existingBooking).Error

	if err == nil {
		return errors.New("mobil tidak tersedia pada tanggal tersebut")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	durasi := booking.TanggalSelesai.Sub(booking.TanggalMulai)
	if durasi < 0 {
		durasi = -durasi
	}
	hours := int(durasi.Hours())
	if hours == 0 {
		hours = 24
	}
	days := float64(hours) / 24.0

	mobil, err := s.MobilRepo.GetMobilByID(booking.MobilID)
	if err != nil {
		return errors.New("gagal mengambil data mobil untuk perhitungan harga")
	}

	booking.TotalHarga = mobil.HargaPerHari * days

	return s.Repo.CreateBooking(booking)
}

func (s *BookingService) UpdateBooking(id uint, booking *models.Booking) error {
	return s.Repo.UpdateBooking(id, booking)
}

func (s *BookingService) DeleteBooking(id uint) error {
	return s.Repo.DeleteBooking(id)
}
