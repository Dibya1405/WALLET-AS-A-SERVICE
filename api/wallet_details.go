package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func WalletDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var wallet entity.Wallet
	result := datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	if result.Error != nil {
		json.NewEncoder(w).Encode("<h1>You have entered invalid key</h1>")
	}
	json.NewEncoder(w).Encode(&wallet)
}
