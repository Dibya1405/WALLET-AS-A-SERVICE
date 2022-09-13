package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func Credit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var wallet entity.Wallet
	amount, _ := strconv.Atoi(params["amount"])
	datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	transaction := entity.Transaction{utils.UUIDGenerator(), wallet.User, wallet.UserId, wallet, wallet.Id, "Credit", amount, time.Now(), time.Now()}
	wallet.Balance = wallet.Balance + amount
	wallet.Updation = time.Now()
	datastore.GetDbHandle().Save(&wallet)
	datastore.GetDbHandle().Create(&transaction)
	json.NewEncoder(w).Encode("Credited successfully!!!")
}
