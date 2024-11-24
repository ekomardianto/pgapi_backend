package config

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var JWT_KEY = []byte("Bismillah9ol4n9$")

type JWTclaim struct {
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Role           string    `json:"role"`
	Name           string    `json:"name"`
	UserId         uuid.UUID `json:"user_id"`
	KodeInstansi   string    `json:"kode_instansi"`
	NamaInstansi   string    `json:"nama_instansi"`
	PerusahaanId   uuid.UUID `json:"per_id"`
	NamaPerusahaan string    `json:"nama_perusahaan"`
	jwt.RegisteredClaims
}
