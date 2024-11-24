package routes

import (
	"eko/api-pg-bpr/controllers/perusahaancontroller"

	"github.com/gorilla/mux"
)

func PerusahaanRoutes(r *mux.Router) {
	r.HandleFunc("/perusahaans", perusahaancontroller.Index).Methods("GET")
	r.HandleFunc("/perusahaans/{id}", perusahaancontroller.GetByIdArray).Methods("GET")
	r.HandleFunc("/perusahaanSearch/{search}", perusahaancontroller.GetSearch).Methods("GET")
	r.HandleFunc("/perusahaan/{id}", perusahaancontroller.Show).Methods("GET")
	r.HandleFunc("/perusahaan", perusahaancontroller.Create).Methods("POST")
	r.HandleFunc("/perusahaan/{id}", perusahaancontroller.Update).Methods("PUT")
	r.HandleFunc("/perusahaan", perusahaancontroller.Delete).Methods("DELETE")
}
