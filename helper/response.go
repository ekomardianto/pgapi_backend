package helper

import (
	"eko/api-pg-bpr/models"
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	respons, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(respons)
}
func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}, message ...string) {
	//buat struktur JSON response
	response := models.ResponseSuccess{
		StatusCode: code,
		Data:       payload,
	}
	// Jika ada nilai untuk message, gunakan yang pertama
	if len(message) > 0 {
		response.Message = message[0]
	} else {
		response.Message = "success" // Nilai default
	}
	ResponseJson(w, http.StatusOK, response)
}
func ResponseSuccessIpay(w http.ResponseWriter, code int, payload interface{}, message ...string) {
	//buat struktur JSON response
	response := models.ResponseSuccessIpay{
		StatusCode: code,
		Success:    true,
		Data:       payload,
	}
	// Jika ada nilai untuk message, gunakan yang pertama
	if len(message) > 0 {
		response.Message = message[0]
	} else {
		response.Message = "Success" // Nilai default
	}
	ResponseJson(w, http.StatusOK, response)
}

func PaginateResponseSuccess(w http.ResponseWriter, code int, payload interface{}, page int, perpage int, count int, countPage int) {
	//buat struktur JSON response
	response := models.PaginateResponseSuccess{
		StatusCode: code,
		Message:    "success",
		Data:       payload,
		Meta: models.Meta{
			PerPage:   perpage,
			Page:      page,
			CountData: count,
			CountPage: countPage,
		},
	}
	ResponseJson(w, http.StatusOK, response)
}
func ResponseFailed(w http.ResponseWriter, code int, message string) {
	//buat struktur JSON response
	response := models.ResponseSuccess{
		StatusCode: code,
		Message:    message,
	}
	ResponseJson(w, http.StatusOK, response)
}
func ResponseFailedIpay(w http.ResponseWriter, code int, message string) {
	//buat struktur JSON response
	response := models.ResponseSuccessIpay{
		StatusCode: code,
		Success:    false,
		Message:    message,
	}
	ResponseJson(w, http.StatusOK, response)
}
func PaginateResponseFailed(w http.ResponseWriter, code int, message string) {
	//buat struktur JSON response
	response := models.PaginateResponseFailed{
		StatusCode: code,
		Message:    message,
		Meta: models.Meta{
			PerPage:   10,
			Page:      1,
			CountData: 0,
			CountPage: 1,
		},
	}
	ResponseJson(w, http.StatusOK, response)
}
func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"message": message})
}
