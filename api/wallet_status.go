package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ChangeStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var wallet entity.Wallet
	datastore.GetDbHandle().First(&wallet, "id=?", params["id"])
	wallet.IsBlocked = !wallet.IsBlocked
	wallet.Updation = time.Now()
	datastore.GetDbHandle().Save(&wallet)
	json.NewEncoder(w).Encode(&wallet)
}
