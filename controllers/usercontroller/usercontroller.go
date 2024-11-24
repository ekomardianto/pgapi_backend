package usercontroller

import (
	"eko/api-pg-bpr/config"
	"eko/api-pg-bpr/helper"
	"eko/api-pg-bpr/models"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ResponseSuccess = helper.ResponseSuccess
var ResponseFailed = helper.ResponseFailed
var PaginateResponseSuccess = helper.PaginateResponseSuccess
var PaginateResponseFailed = helper.PaginateResponseFailed
var mainTable = "users"

// Service Default ***********************
func Index(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := models.DB.Preload("Instansi").Preload("Instansi.Perusahaan").Find(&users).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, users)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User // type
	if err := models.DB.First(&user, "id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "User tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseSuccess(w, http.StatusOK, user)
}
func Create(w http.ResponseWriter, r *http.Request) {

	var inputUser models.RegisterUser // type
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

	// Cek apakah username sudah ada di database
	var existingUsername models.User
	if err := models.DB.Where("username = ?", inputUser.Username).First(&existingUsername).Error; err == nil {
		// Jika ditemukan username yang sama, return error
		ResponseFailed(w, http.StatusBadRequest, "username sudah ada! gunakan username lain.")
		return
	} else if err != gorm.ErrRecordNotFound {
		// Jika ada error selain tidak ditemukan, return error
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	// jika tidak ditemukan duplikat username maka simpan username baru ke database
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

	var response = map[string]string{"username": dataUser.Username}
	ResponseSuccess(w, http.StatusOK, response, "User Berhasil ditambahkan")
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	var user models.User // type
	if models.DB.Delete(&user, "id = ?", input["id"]).RowsAffected == 0 {
		ResponseFailed(w, http.StatusInternalServerError, "tidak dapat menghapus User")
		return
	}
	var response = map[string]string{"message": "User berhasil dihapus", "id": input["id"]}
	ResponseSuccess(w, http.StatusOK, response)
}

// End Service Default *******************

// Service Untuk Profile Management **********************************************
func Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User // type
	var userProfile models.UserProfile

	fmt.Printf("username: %v\n", username)

	if err := models.DB.Find(&user, "username = ?", username).Scan(&userProfile).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, userProfile)
}
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	//bandingkan username dalam token dengan username yg akan diubah
	usernameInfo := r.Context().Value("userInfo").(*config.JWTclaim).Username

	var usernameData string
	models.DB.Raw("SELECT username FROM users WHERE id = ?", id).Scan(&usernameData)

	if usernameInfo == usernameData {
		ResponseFailed(w, http.StatusInternalServerError, "tidak dapat mengupdate role sendiri")
		return
	}

	//TAMBAHKAN WAKTU UPDATED_AT KE DATABASE
	user.UpdatedAt = time.Now().UnixMilli()

	if err := models.DB.Exec("UPDATE users SET role_id = ?, updated_at = ? WHERE id = ?",
		user.RoleId, user.UpdatedAt, id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := uuid.Parse(id)

	if err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	user.Id = userId

	var response = map[string]string{"username": usernameData, "role_id": user.RoleId}
	ResponseSuccess(w, http.StatusOK, response)
}
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var userChangePassword models.ChangePassword // Struct untuk input data
	decoder := json.NewDecoder(r.Body)

	// Decode JSON body
	if err := decoder.Decode(&userChangePassword); err != nil {
		ResponseFailed(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	// Ambil username dan password dari database berdasarkan ID
	var usernamePwd models.UsernamePwd
	if err := models.DB.Raw("SELECT username, password FROM users WHERE id = ?", id).Scan(&usernamePwd).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, "Failed to fetch user data")
		return
	}

	// Periksa apakah pengguna mencoba mengubah password miliknya sendiri
	usernameInfo := r.Context().Value("userInfo").(*config.JWTclaim).Username
	if usernameInfo == usernamePwd.Username {
		ResponseFailed(w, http.StatusForbidden, "Cannot change your own password")
		return
	}

	// Periksa apakah password baru dan konfirmasi password cocok
	if userChangePassword.NewPassword != userChangePassword.ConfirmPassword {
		ResponseFailed(w, http.StatusBadRequest, "New password and confirm password do not match")
		return
	}

	// Verifikasi password lama dengan yang ada di database
	if err := bcrypt.CompareHashAndPassword([]byte(usernamePwd.Password), []byte(userChangePassword.OldPassword)); err != nil {
		ResponseFailed(w, http.StatusUnauthorized, "Old password is incorrect")
		return
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userChangePassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ResponseFailed(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Update password di database
	if err := models.DB.Exec("UPDATE users SET password = ?, updated_at = ? WHERE id = ?",
		string(hashedPassword), time.Now().UnixMilli(), id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, "Failed to update password")
		return
	}

	// Berikan respons sukses
	ResponseSuccess(w, http.StatusOK, map[string]string{
		"message": "Password updated successfully",
	})
}
func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UpdateUser // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	//TAMBAHKAN WAKTU UPDATED_AT KE DATABASE
	user.UpdatedAt = time.Now().UnixMilli()
	if err := models.DB.Exec("UPDATE users SET nama = ?, email = ?, instansi_id = ?, status = ?, updated_at = ? WHERE id = ?",
		user.Nama, user.Email, user.InstansiId, user.Status, user.UpdatedAt, id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Ambil username berdasarkan ID dari database
	var username string
	if err := models.DB.Raw("SELECT username FROM users WHERE id = ?", id).Scan(&username).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, "Gagal mengambil username")
		return
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	user.Id = userId
	var response = map[string]string{"username": username}
	ResponseSuccess(w, http.StatusOK, response, "Data User Berhasil diubah")
}

// END User Profile Management ****************************************************

// Service Untuk Pagination *******************************************************
func IndexCustomsPaginate(w http.ResponseWriter, r *http.Request) {
	// ambil variable prpage dari URL
	perpage := r.URL.Query().Get("perpage")
	if perpage == "" {
		perpage = "10"
	}
	if perpage > "50" {
		perpage = "50"
	}
	perpageInt, _ := strconv.Atoi(perpage)

	// hitung jumlah page dari db
	var count = int64(0)
	models.DB.Table(mainTable).Count(&count)
	var countPage = int(math.Ceil(float64(count) / float64(perpageInt)))

	// ambil variable Page dari URL
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)
	if pageInt > countPage {
		pageInt = countPage
	}

	var users []models.User
	// Join products and categories tables and select product data along with category name
	if err := models.DB.
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&users).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) == 0 {
		PaginateResponseSuccess(w, http.StatusOK, []interface{}{}, 1, 10, 1, countPage)
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, users, pageInt, perpageInt, int(count), countPage)
}
func SearchPaginate(w http.ResponseWriter, r *http.Request) {
	// ambil variable SerachParameter dari URL
	vars := mux.Vars(r)
	searchParam := vars["searchParam"]

	// ambil variable perpage dari URL
	perpage := r.URL.Query().Get("perpage")
	if perpage == "" {
		perpage = "3"
	}
	if perpage > "50" {
		perpage = "50"
	}
	perpageInt, _ := strconv.Atoi(perpage)

	// deklarasikan variable countPage dari db
	var count = int64(0)
	models.DB.Table(mainTable).
		Where("users.nama LIKE ?", "%"+searchParam+"%").
		Or("users.username LIKE ?", "%"+searchParam+"%").
		Or("users.email LIKE ?", "%"+searchParam+"%").
		Count(&count)
	var countPage = int(math.Ceil(float64(count) / float64(perpageInt)))

	// ambil variable Page dari URL
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)
	if pageInt > countPage {
		pageInt = countPage
	}

	var users []models.User // type
	if err := models.DB.
		Where("users.nama LIKE ? OR users.username LIKE ? OR users.email LIKE ?", "%"+searchParam+"%", "%"+searchParam+"%", "%"+searchParam+"%").
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&users).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) == 0 {
		ResponseSuccess(w, http.StatusOK, []interface{}{})
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, users, pageInt, perpageInt, int(count), countPage)
}

// END Service Untuk Pagination ****************************************************
