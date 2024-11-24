package helper

import (
	"eko/api-pg-bpr/config"
	"eko/api-pg-bpr/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(dataUser *models.User) (string, error) {
	expTime := time.Now().Add(time.Minute * 120)
	claims := &config.JWTclaim{
		Username:       dataUser.Username,
		Email:          dataUser.Email,
		Role:           dataUser.RoleId,
		Name:           dataUser.Nama,
		UserId:         dataUser.Id,
		KodeInstansi:   string(dataUser.Instansi.KodeInstansi),
		NamaInstansi:   dataUser.Instansi.Name,
		PerusahaanId:   dataUser.Instansi.PerusahaanId,
		NamaPerusahaan: dataUser.Instansi.Perusahaan.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "api-pg-bpr",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgoritm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgoritm.SignedString(config.JWT_KEY)
	return token, err
}

func ValidateToken(tokenString string) (*config.JWTclaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.JWTclaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unautorized, token invalid!")
	}

	claim, ok := token.Claims.(*config.JWTclaim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unautorized, token invalid!")
	}

	return claim, nil

}
