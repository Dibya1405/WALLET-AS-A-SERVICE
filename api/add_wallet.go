package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func AddOneWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	logger.GetMyLogger().Warn("Getting the user from the user table")
	datastore.GetDbHandle().First(&user, "id=?", params["id"])
	var wallet entity.Wallet
	wallet.Id = utils.UUIDGenerator()
	wallet.User = user
	wallet.Creation = time.Now()
	wallet.Updation = time.Now()
	wallet.UserId = params["id"]
	wallet.IsDeleted = false
	wallet.Balance = 0

	logger.GetMyLogger().Warn("Taking the values from request body....")
	json.NewDecoder(r.Body).Decode(&wallet)
	datastore.GetDbHandle().Create(&wallet)
	logger.GetMyLogger().Info("New Wallet Created")
	json.NewEncoder(w).Encode(&wallet)
}
