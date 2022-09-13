package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var wallet entity.Wallet
	var transactions []entity.Transaction
	//datastore.GetDbHandle().Delete(&wallet, "id=?", params["id"])
	datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	wallet.IsDeleted = true
	wallet.Updation = time.Now()
	datastore.GetDbHandle().Save(&wallet)
	datastore.GetDbHandle().Where("wallet_Id=?", params["id"]).Delete(&transactions)

	json.NewEncoder(w).Encode("Deleted Successfully!!!!")
}
