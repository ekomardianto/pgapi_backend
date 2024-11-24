package authcontroller

import (
	"eko/api-pg-bpr/helper"
	"eko/api-pg-bpr/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ResponseFailed = helper.ResponseFailed
var ResponseSuccess = helper.ResponseSuccess

func Login(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("masuk login")
	var inputUser models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&inputUser); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var dataUser models.User
	if err := models.DB.Preload("Instansi").Preload("Instansi.Perusahaan").Where("username = ?", inputUser.Username).Find(&dataUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if dataUser.Username != inputUser.Username {
		ResponseFailed(w, 402, "username salah")
		return
	} else {

		if err := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(inputUser.Password)); err != nil {
			ResponseFailed(w, http.StatusUnauthorized, "password salah")
			return
		}

		//Proses Pembuatan token
		token, err := helper.GenerateToken(&dataUser)
		if err != nil {
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
		// ResponseSuccess(w, http.StatusOK, dataUser)
		ResponseSuccess(w, http.StatusOK, map[string]string{
			"token": token,
			// "username": dataUser.Username,
			// "role":     dataUser.RoleId,
			// "email":    dataUser.Email,
			// "name":     dataUser.Nama,
		})
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	var inputUser models.RegisterUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputUser); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Konversi CreateUserRequest ke models.User
	user := models.User{
		Id:         uuid.New(),
		Nama:       inputUser.Nama,
		Email:      inputUser.Email,
		Username:   inputUser.Username,
		Password:   inputUser.Password,
		Phone:      inputUser.Phone,
		RoleId:     inputUser.RoleId,
		InstansiId: inputUser.InstansiId,
		Status:     inputUser.Status,
		CreatedAt:  time.Now().UnixMilli(),
		UpdatedAt:  time.Now().UnixMilli(),
	}
	hasingPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hasingPassword)
	user.Status = "1"
	user.RoleId = "member"

	if err := models.DB.Create(&user).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	var dataUser models.User
	if err := models.DB.Preload("Instansi").Preload("Instansi.Perusahaan").Where("id = ?", user.Id).Find(&dataUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseSuccess(w, http.StatusOK, dataUser)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	//hapus tokenn dari whitlist

	ResponseSuccess(w, http.StatusOK, "logout berhasil")
}
