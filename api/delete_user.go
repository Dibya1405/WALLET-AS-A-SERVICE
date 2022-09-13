package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user entity.User
	var wallets []entity.Wallet
	//var transactions []entity.Transaction
	//datastore.GetDbHandle().Delete(&user, "id=?", params["id"])
	datastore.GetDbHandle().First(&user, "id=?", params["id"])
	user.IsDeleted = true
	user.Updation = time.Now()
	datastore.GetDbHandle().Save(&user)
	datastore.GetDbHandle().Where("user_Id=?", params["id"]).Delete(&wallets)
	//datastore.GetDbHandle().Where("user_Id=?", params["id"]).Delete(&transactions)
	json.NewEncoder(w).Encode("Deleted Successfully!!!!")
}
