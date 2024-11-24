package routes

import (
	"eko/api-pg-bpr/controllers/authcontroller"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	// router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
}
