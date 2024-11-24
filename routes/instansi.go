package routes

import (
	"eko/api-pg-bpr/controllers/instansicontroller"

	"github.com/gorilla/mux"
)

func InstansiRoutes(r *mux.Router) {
	r.HandleFunc("/instansis", instansicontroller.Index).Methods("GET")
	r.HandleFunc("/instansi/{id}", instansicontroller.Show).Methods("GET")
	r.HandleFunc("/instansisPerusahaan/{per_id}", instansicontroller.IndexByPer).Methods("GET")
	r.HandleFunc("/instansi", instansicontroller.Create).Methods("POST")
	r.HandleFunc("/instansi/{id}", instansicontroller.Update).Methods("PUT")
	r.HandleFunc("/instansi", instansicontroller.Delete).Methods("DELETE")
}
