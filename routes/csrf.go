package routes

import (
	"eko/api-pg-bpr/controllers/csrfcontroller"

	"github.com/gorilla/mux"
)

func CSRFRoutes(r *mux.Router) {
	r.HandleFunc("/csrf", csrfcontroller.GenerateCSRFToken).Methods("GET")
}
