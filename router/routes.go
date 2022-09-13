package router

import (
	"WALLET-AS-A-SERVICE/api"

	"github.com/gorilla/mux"
)

func InitializeRouter(r *mux.Router) {
	r.HandleFunc("/signup", api.Signup).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/profile/{id}", api.GetProfile).Methods("GET")
	r.HandleFunc("/deleteUser/{id}", api.DeleteOneUser).Methods("PATCH")
	r.HandleFunc("/walletsAll/{id}", api.GetAllWallets).Methods("GET")
	r.HandleFunc("/changeStatus/{id}", api.ChangeStatus).Methods("PATCH")
	r.HandleFunc("/addWallet/{id}", api.AddOneWallet).Methods("POST")
	r.HandleFunc("/deleteWallet/{id}", api.DeleteWallet).Methods("PATCH")
	r.HandleFunc("/walletDetails/{id}", api.WalletDetails).Methods("GET")
	r.HandleFunc("/wallet/credit/{amount}/{id}", api.Credit).Methods("PATCH")
	r.HandleFunc("/wallet/debit/{amount}/{id}", api.Debit).Methods("PATCH")
	r.HandleFunc("/history/{id}", api.History).Methods("GET")
}
