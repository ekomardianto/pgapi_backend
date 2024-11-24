package roleusercontroller

import (
	"eko/api-pg-bpr/helper"
	"eko/api-pg-bpr/models"
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseSuccess = helper.ResponseSuccess
var ResponseFailed = helper.ResponseFailed
var PaginateResponseSuccess = helper.PaginateResponseSuccess
var PaginateResponseFailed = helper.PaginateResponseFailed
var mainTable = "role_users"

// default Service **********************************************
func Index(w http.ResponseWriter, r *http.Request) {
	var roles []models.RoleUser
	if err := models.DB.Find(&roles).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, roles)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var role models.RoleUser // type
	if err := models.DB.First(&role, "id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "Role User tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseSuccess(w, http.StatusOK, role)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var role models.RoleUser // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&role); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&role).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response = map[string]string{"role": role.Role}
	ResponseSuccess(w, http.StatusOK, response)
}
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var role models.RoleUser // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&role); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Exec("UPDATE role_users SET role = ? WHERE id = ?",
		role.Role, id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	roleId, err := uuid.Parse(id)
	if err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	role.Id = roleId
	var response = map[string]string{"role": role.Role}
	ResponseSuccess(w, http.StatusOK, response)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	var role models.RoleUser // type
	if models.DB.Delete(&role, "id = ?", input["id"]).RowsAffected == 0 {
		ResponseFailed(w, http.StatusInternalServerError, "tidak dapat menghapus Role User")
		return
	}
	var response = map[string]string{"message": "Role User berhasil dihapus", "id": input["id"]}
	ResponseSuccess(w, http.StatusOK, response)
}

// Service Pagination ********************************************
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

	var roleUser []models.RoleUser
	// Join products and categories tables and select product data along with category name
	if err := models.DB.
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&roleUser).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(roleUser) == 0 {
		PaginateResponseSuccess(w, http.StatusOK, []interface{}{}, 1, 10, 1, countPage)
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, roleUser, pageInt, perpageInt, int(count), countPage)
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
		Where("role_users.role LIKE ?", "%"+searchParam+"%").
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

	var roleUsers []models.RoleUser // type
	if err := models.DB.
		Where("role_users.role LIKE ?", "%"+searchParam+"%").
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&roleUsers).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(roleUsers) == 0 {
		ResponseSuccess(w, http.StatusOK, []interface{}{})
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, roleUsers, pageInt, perpageInt, int(count), countPage)
}
