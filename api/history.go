package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var transactions []entity.Transaction

	logger.GetMyLogger().Warn("Fetching all the transaction history")
	datastore.GetDbHandle().Find(&transactions, "wallet_id=?", params["id"])
	json.NewEncoder(w).Encode(&transactions)
}
