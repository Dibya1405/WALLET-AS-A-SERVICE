package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"encoding/json"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []entity.User
	logger.GetMyLogger().Warn("Getting all users...")
	datastore.GetDbHandle().Find(&users)
	json.NewEncoder(w).Encode(&users)
}
