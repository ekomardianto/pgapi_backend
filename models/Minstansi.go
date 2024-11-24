package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Instansi struct {
	Id           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KodeInstansi string     `gorm:"type:varchar(10);not null" json:"kode_instansi"`
	PerusahaanId uuid.UUID  `gorm:"type:char(36);not null" json:"-"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	Alamat       string     `gorm:"type:varchar(255);not null" json:"alamat"`
	Telp         string     `gorm:"type:varchar(255);not null" json:"phone"`
	CreatedAt    int64      `gorm:"type:bigint;not null" json:"created_at"`
	UpdatedAt    int64      `gorm:"type:bigint;not null" json:"updated_at"`
	Perusahaan   Perusahaan `gorm:"foreignKey:PerusahaanId" json:"perusahaan"`
}
type AddInstansi struct {
	Id           uuid.UUID `json:"id"`
	KodeInstansi string    `json:"kode_instansi"`
	PerusahaanId uuid.UUID `json:"perusahaan_id"`
	Name         string    `json:"name"`
	Alamat       string    `json:"alamat"`
	Telp         string    `json:"phone"`
	CreatedAt    int64     `json:"created_at"`
	UpdatedAt    int64     `json:"updated_at"`
}
type InstansiPerusahaan struct {
	Instansi
	NamaPerusahaan string `json:"name_per"`
}

func (j *Instansi) BeforeCreate(tx *gorm.DB) (err error) {
	j.Id = uuid.New()
	return
}
