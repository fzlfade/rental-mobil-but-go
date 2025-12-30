package models

import (
	"time"
)

type Booking struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	MobilID        uint      `json:"mobil_id" gorm:"not null"`
	TanggalMulai   time.Time `json:"tanggal_mulai" gorm:"not null"`
	TanggalSelesai time.Time `json:"tanggal_selesai" gorm:"not null"`
	TotalHarga     float64   `json:"total_harga" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	User  Mobil `json:"user" gorm:"foreignKey:UserID"`
	Mobil Mobil `json:"mobil" gorm:"foreignKey:MobilID"`
}
