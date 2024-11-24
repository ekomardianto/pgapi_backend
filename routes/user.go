package routes

import (
	"eko/api-pg-bpr/controllers/usercontroller"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	// default routes
	r.HandleFunc("/users", usercontroller.Index).Methods("GET")
	r.HandleFunc("/user/{id}", usercontroller.Show).Methods("GET")
	r.HandleFunc("/user", usercontroller.Create).Methods("POST")
	r.HandleFunc("/userprofile/{username}", usercontroller.Profile).Methods("GET")
	r.HandleFunc("/user/roleupdate/{id}", usercontroller.UpdateRole).Methods("PUT")
	r.HandleFunc("/user/dataupdate/{id}", usercontroller.UpdateUserData).Methods("PUT")
	r.HandleFunc("/user/changepassword/{id}", usercontroller.UpdateUserPassword).Methods("PUT")
	r.HandleFunc("/user", usercontroller.DeleteUser).Methods("DELETE")

	// custom routes Pagination
	r.HandleFunc("/usersPaginate", usercontroller.IndexCustomsPaginate).Methods("GET")
	r.HandleFunc("/userSearch/{searchParam}", usercontroller.SearchPaginate).Methods("GET")
}
