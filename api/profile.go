package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user entity.User
	logger.GetMyLogger().Warn("Fetching user info..")
	result := datastore.GetDbHandle().First(&user, "id=?", params["id"])
	if result.Error != nil {
		logger.GetMyLogger().Error("Invalid key...")
		json.NewEncoder(w).Encode("<h1>You have entered invalid key</h1>")
	} else {
		logger.GetMyLogger().Info("Data fetched successfully!!!")
		json.NewEncoder(w).Encode(&user)
	}
}
