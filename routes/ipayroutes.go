package routes

import (
	"eko/api-pg-bpr/controllers/ipaycontroller"

	"github.com/gorilla/mux"
)

func Ipayroutes(r *mux.Router) {
	r.HandleFunc("/ipayinquirybalance", ipaycontroller.InquiryBalance).Methods("POST")
	r.HandleFunc("/ipaychecktransaction", ipaycontroller.CheckTransaction).Methods("POST")
	r.HandleFunc("/ipayhistorytransaction", ipaycontroller.HistoryTransaction).Methods("POST")

	//transaksi
	r.HandleFunc("/ipaycreatetransaction", ipaycontroller.CreateTransaction).Methods("POST")
}
