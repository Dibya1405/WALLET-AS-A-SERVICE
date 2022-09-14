package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func Debit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var wallet entity.Wallet
	logger.GetMyLogger().Info("Getting the amount from the user....")
	amount, _ := strconv.Atoi(params["amount"])
	datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	transaction := entity.Transaction{utils.UUIDGenerator(), wallet.User, wallet.UserId, wallet, wallet.Id, "Debit", amount, time.Now(), time.Now()}
	wallet.Balance = wallet.Balance - amount
	wallet.Updation = time.Now()
	datastore.GetDbHandle().Save(&wallet)
	logger.GetMyLogger().Info("Wallet Balance updated...")
	datastore.GetDbHandle().Create(&transaction)
	logger.GetMyLogger().Warn("Transaction of type debit added into the table ")
	json.NewEncoder(w).Encode("Debited Successfully!!!")
}
