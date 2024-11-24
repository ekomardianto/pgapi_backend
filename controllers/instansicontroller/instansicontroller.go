package instansicontroller

import (
	"eko/api-pg-bpr/helper"
	"eko/api-pg-bpr/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseSuccess = helper.ResponseSuccess
var ResponseFailed = helper.ResponseFailed

// Service Default **********************************
func Index(w http.ResponseWriter, r *http.Request) {
	var instansis []models.Instansi
	if err := models.DB.Order("name asc").Preload("Perusahaan").Find(&instansis).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, instansis)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var instansi models.Instansi // type
	if err := models.DB.First(&instansi, "id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseFailed(w, http.StatusNotFound, "Instansi tidak ditemukan")
			return
		default:
			ResponseFailed(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseSuccess(w, http.StatusOK, instansi)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var inputReqInstansi models.AddInstansi // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&inputReqInstansi); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// deklarasikan dan masukkan ke model Instansi
	instansi := models.Instansi{
		Id:           uuid.New(),
		KodeInstansi: inputReqInstansi.KodeInstansi,
		PerusahaanId: inputReqInstansi.PerusahaanId,
		Name:         inputReqInstansi.Name,
		Alamat:       inputReqInstansi.Alamat,
		Telp:         inputReqInstansi.Telp,
		CreatedAt:    time.Now().UnixMilli(),
		UpdatedAt:    time.Now().UnixMilli(),
	}

	// Cek apakah kode_instansi sudah ada di database
	var existingInstansi models.Instansi
	if err := models.DB.Where("kode_instansi = ?", instansi.KodeInstansi).First(&existingInstansi).Error; err == nil {
		// Jika ditemukan kode_instansi yang sama, return error
		ResponseFailed(w, http.StatusBadRequest, "Kode instansi sudah ada! gunakan kode lain.")
		return
	} else if err != gorm.ErrRecordNotFound {
		// Jika ada error selain tidak ditemukan, return error
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	//add Instansi
	if err := models.DB.Create(&instansi).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response = map[string]string{"name": instansi.Name}
	ResponseSuccess(w, http.StatusOK, response, "Instansi Berhasil ditambahkan")
}
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var instansi models.Instansi // type
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&instansi); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Cek apakah kode_instansi sudah ada di database
	var existingInstansi models.Instansi
	if err := models.DB.Where("kode_instansi = ?", instansi.KodeInstansi).First(&existingInstansi).Error; err == nil {
		// Jika ditemukan kode_instansi yang sama, return error
		ResponseFailed(w, http.StatusBadRequest, "Kode instansi sudah ada! gunakan kode lain.")
		return
	} else if err != gorm.ErrRecordNotFound {
		// Jika ada error selain tidak ditemukan, return error
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	//TAMBAHKAN WAKTU UPDATED_AT KE DATABASE
	instansi.UpdatedAt = time.Now().UnixMilli()
	if err := models.DB.Exec("UPDATE instansis SET name = ?, alamat = ?,telp = ?, kode_instansi = ?, perusahaan_id = ?, updated_at = ? WHERE id = ?",
		instansi.Name, instansi.Alamat, instansi.Telp, instansi.KodeInstansi, instansi.Perusahaan.Id, instansi.UpdatedAt, id).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}
	instansiId, err := uuid.Parse(id)
	if err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	instansi.Id = instansiId
	var response = map[string]string{"name": instansi.Name}
	ResponseSuccess(w, http.StatusOK, response, "Instansi Berhasil diubah")
}
func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	var instansi models.Instansi // type
	if models.DB.Delete(&instansi, "id = ?", input["id"]).RowsAffected == 0 {
		ResponseFailed(w, http.StatusInternalServerError, "tidak dapat menghapus Instansi")
		return
	}
	var response = map[string]string{"message": "Instansi berhasil dihapus", "id": input["id"]}
	ResponseSuccess(w, http.StatusOK, response)
}

// END Service Default ********************************
// Start Service BY Per_ID ********************************
func IndexByPer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	perId := vars["per_id"]

	var instansis []models.Instansi
	if err := models.DB.Where("perusahaan_id = ?", perId).Preload("Perusahaan").
		Find(&instansis).Error; err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, instansis)
}
