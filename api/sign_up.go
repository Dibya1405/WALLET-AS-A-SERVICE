package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	logger.GetMyLogger().Warn("Taking the values from request body....")
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = utils.UUIDGenerator()
	user.Creation = time.Now()
	user.Updation = time.Now()
	user.IsDeleted = false
	datastore.GetDbHandle().Create(&user)
	logger.GetMyLogger().Info("New User Created")
	json.NewEncoder(w).Encode(&user)
}
