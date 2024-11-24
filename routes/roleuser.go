package routes

import (
	"eko/api-pg-bpr/controllers/roleusercontroller"

	"github.com/gorilla/mux"
)

func RoleUserRoutes(r *mux.Router) {
	// Routing Default
	r.HandleFunc("/roles", roleusercontroller.Index).Methods("GET")
	r.HandleFunc("/role/{id}", roleusercontroller.Show).Methods("GET")
	r.HandleFunc("/role", roleusercontroller.Create).Methods("POST")
	r.HandleFunc("/role/{id}", roleusercontroller.Update).Methods("PUT")
	r.HandleFunc("/role", roleusercontroller.Delete).Methods("DELETE")

	// Routing Custom Paginatrion
	r.HandleFunc("/rolesPaginate", roleusercontroller.IndexCustomsPaginate).Methods("GET")
	r.HandleFunc("/roleSearch/{searchParam}", roleusercontroller.SearchPaginate).Methods("GET")
}
