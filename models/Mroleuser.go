package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleUser struct {
	Id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Role string    `gorm:"type:varchar(36);not null" json:"role"`
}

// Fungsi BeforeCreate untuk menggenerate UUID sebelum record disimpan
func (j *RoleUser) BeforeCreate(tx *gorm.DB) (err error) {
	j.Id = uuid.New()
	return
}
