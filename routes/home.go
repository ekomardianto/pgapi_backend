package routes

import (
	"eko/api-pg-bpr/controllers/homecontroller"

	"github.com/gorilla/mux"
)

func HomeRoutes(r *mux.Router) {
	r.HandleFunc("/", homecontroller.Index).Methods("GET")
}
