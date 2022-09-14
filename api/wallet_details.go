package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func WalletDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var wallet entity.Wallet
	logger.GetMyLogger().Warn("Fetching wallet info..")
	result := datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	if result.Error != nil {
		logger.GetMyLogger().Error("Invalid key...")
		json.NewEncoder(w).Encode("<h1>You have entered invalid key</h1>")
	} else {
		logger.GetMyLogger().Info("Data fetched successfully!!!")
		json.NewEncoder(w).Encode(&wallet)
	}
}
