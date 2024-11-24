package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Nama       string    `gorm:"type:varchar(255);not null" json:"nama"`
	Email      string    `gorm:"type:varchar(255);not null" json:"email"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	Phone      string    `gorm:"type:varchar(255);not null" json:"phone"`
	RoleId     string    `gorm:"type:varchar(36);not null" json:"role_id"`
	InstansiId uuid.UUID `gorm:"type:varchar(36);not null" json:"-"`
	Status     string    `gorm:"type:varchar(2);not null" json:"status"`
	CreatedAt  int64     `gorm:"type:bigint;not null" json:"created_at"`
	UpdatedAt  int64     `gorm:"type:bigint;not null" json:"updated_at"`
	Instansi   Instansi  `gorm:"foreignkey:InstansiId" json:"instansi"`
}

type RegisterUser struct {
	Nama       string    `json:"nama"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	RoleId     string    `json:"role_id"`
	InstansiId uuid.UUID `json:"instansi_id"`
	Status     string    `json:"status"`
	CreatedAt  int64     `json:"created_at"`
	UpdatedAt  int64     `json:"updated_at"`
}
type UserProfile struct {
	Id         uuid.UUID `json:"id"`
	Nama       string    `json:"nama"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	RoleId     string    `json:"role_id"`
	InstansiId uuid.UUID `json:"instansi_id"`
	CreatedAt  int64     `json:"created_at"`
	UpdatedAt  int64     `json:"updated_at"`
	Status     string    `json:"status"`
}
type UpdateUser struct {
	Id         uuid.UUID `json:"id"`
	Nama       string    `json:"nama"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	RoleId     string    `json:"role_id"`
	InstansiId uuid.UUID `json:"instansi_id"`
	Status     string    `json:"status"`
	UpdatedAt  int64     `json:"updated_at"`
}
type UsernamePwd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	UpdatedAt       int64  `json:"updated_at"`
}
