package perusahaancontroller

import (
	"eko/api-pg-bpr/helper"
	"eko/api-pg-bpr/models"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseSuccess = helper.ResponseSuccess
var ResponseFailed = helper.ResponseFailed
var PaginateResponseSuccess = helper.PaginateResponseSuccess
var PaginateResponseFailed = helper.PaginateResponseFailed
var mainTable = "perusahaans"

// Service Default ***********************
func Index(w http.ResponseWriter, r *http.Request) {
	var perusahaan []models.Perusahaan
	if err := models.DB.Find(&perusahaan).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(w, http.StatusOK, perusahaan)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perusahaan models.Perusahaan // type
	if err := models.DB.First(&perusahaan, "id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "perusahaan tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseSuccess(w, http.StatusOK, perusahaan)
}
func GetByIdArray(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perusahaans []models.Perusahaan // type
	if err := models.DB.Where("id = ?", id).Find(&perusahaans).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "perusahaan tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseSuccess(w, http.StatusOK, perusahaans)
}
func GetSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchParam := vars["search"]
	var perusahaans []models.Perusahaan // type
	if err := models.DB.
		Where("name ILIKE ? OR alamat ILIKE ? OR kode_perusahaan ILIKE ?", "%"+searchParam+"%", "%"+searchParam+"%", "%"+searchParam+"%").
		Find(&perusahaans).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "perusahaan tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if len(perusahaans) == 0 {
		ResponseFailed(w, 201, "perusahaans tidak ditemukan")
		return
	}
	ResponseSuccess(w, http.StatusOK, perusahaans)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var perusahaan models.Perusahaan // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&perusahaan); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	perusahaan.CreatedAt = time.Now().UnixMilli()
	perusahaan.UpdatedAt = time.Now().UnixMilli()
	if err := models.DB.Create(&perusahaan).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(w, http.StatusOK, perusahaan)
}
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perusahaan models.Perusahaan // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&perusahaan); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	perusahaan.UpdatedAt = time.Now().UnixMilli()
	if err := models.DB.Exec("UPDATE perusahaans SET kode_perusahaan = ?, name = ?, alamat = ?, telp = ?, updated_at = ? WHERE id = ?",
		perusahaan.KodePerusahaan, perusahaan.Name, perusahaan.Alamat, perusahaan.Telp, perusahaan.UpdatedAt, id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	perusahaanId, err := uuid.Parse(id)
	if err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	perusahaan.Id = perusahaanId

	ResponseSuccess(w, http.StatusOK, perusahaan)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	var perusahaan models.Perusahaan // type
	if models.DB.Delete(&perusahaan, "id = ?", input["id"]).RowsAffected == 0 {
		ResponseFailed(w, http.StatusInternalServerError, "tidak dapat menghapus Perusahaan")
		return
	}
	var response = map[string]string{"message": "Perusahaan berhasil dihapus", "id": input["id"]}
	ResponseSuccess(w, http.StatusOK, response)
}

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

	var perusahaan []models.Perusahaan
	// Join products and categories tables and select product data along with category name
	if err := models.DB.
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&perusahaan).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(perusahaan) == 0 {
		PaginateResponseSuccess(w, http.StatusOK, []interface{}{}, 1, 10, 1, countPage)
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, perusahaan, pageInt, perpageInt, int(count), countPage)
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
		Where("perusahaans.name LIKE ?", "%"+searchParam+"%").
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

	var perusahaans []models.Perusahaan // type
	if err := models.DB.
		Where("perusahaans.name LIKE ?", "%"+searchParam+"%").
		Limit(perpageInt).Offset((pageInt - 1) * perpageInt).
		Find(&perusahaans).Error; err != nil {
		PaginateResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(perusahaans) == 0 {
		ResponseSuccess(w, http.StatusOK, []interface{}{})
		return
	}
	PaginateResponseSuccess(w, http.StatusOK, perusahaans, pageInt, perpageInt, int(count), countPage)
}

// END Service Untuk Pagination ****************************************************
