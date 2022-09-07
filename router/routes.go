package router

import (
	"WALLET-AS-A-SERVICE/api"

	"github.com/gorilla/mux"
)

func InitializeRouter(r *mux.Router) {
	r.HandleFunc("/signup", api.DemoFunc).Methods("POST")
	r.HandleFunc("/users", api.DemoFunc).Methods("GET")
	r.HandleFunc("/profile/{id}", api.DemoFunc).Methods("GET")
	r.HandleFunc("/deleteUser", api.DemoFunc).Methods("DELETE")
	r.HandleFunc("/walletsAll/{id}", api.DemoFunc).Methods("GET")
	r.HandleFunc("/changeStatus", api.DemoFunc).Methods("PATCH")
	r.HandleFunc("/addWallet", api.DemoFunc).Methods("POST")
	r.HandleFunc("/deleteWallet", api.DemoFunc).Methods("DELETE")
	r.HandleFunc("/walletDetails/{id}", api.DemoFunc).Methods("GET")
	r.HandleFunc("/wallet/credit", api.DemoFunc).Methods("PUT")
	r.HandleFunc("/wallet/debit", api.DemoFunc).Methods("PUT")
	r.HandleFunc("/history/{id}", api.DemoFunc).Methods("GET")
}
