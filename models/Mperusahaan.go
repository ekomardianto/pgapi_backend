package models

import (
	"github.com/google/uuid"
)

type Perusahaan struct {
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KodePerusahaan string    `gorm:"type:varchar(10);not null" json:"kode_per"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Alamat         string    `gorm:"type:varchar(255);not null" json:"alamat"`
	Telp           string    `gorm:"type:varchar(255);not null" json:"phone"`
	CreatedAt      int64     `gorm:"type:bigint;not null" json:"created_at"`
	UpdatedAt      int64     `gorm:"type:bigint;not null" json:"updated_at"`
	// Instansi       []Instansi `gorm:"foreignKey:PerusahaansId" json:"perusahaan"`
}
type PerusahaanUpdate struct {
	KodePerusahaan *string `json:"kode_per,omitempty"`
	Name           *string `json:"name,omitempty"`
	Alamat         *string `json:"alamat,omitempty"`
	Telp           *string `json:"phone,omitempty"`
	UpdatedAt      *int64  `json:"updated_at,omitempty"`
}
