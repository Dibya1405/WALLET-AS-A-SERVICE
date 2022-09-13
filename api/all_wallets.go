package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllWallets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var wallets []entity.Wallet
	datastore.GetDbHandle().Where("user_Id=?", params["id"]).Find(&wallets)
	json.NewEncoder(w).Encode(&wallets)
}
