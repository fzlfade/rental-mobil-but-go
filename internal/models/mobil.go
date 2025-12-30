package models

import (
	"time"
)

type Mobil struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Nama         string    `json:"nama" gorm:"not null"`
	Merek        string    `json:"merek" gorm:"not null"`
	Tahun        int       `json:"tahun" gorm:"not null"`
	Plat         string    `json:"plat" gorm:"type:varchar(20);uniqueIndex;not null"`
	HargaPerHari float64   `json:"harga_per_hari" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
