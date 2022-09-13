package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = utils.UUIDGenerator()
	user.Creation = time.Now()
	user.Updation = time.Now()
	user.IsDeleted = false
	datastore.GetDbHandle().Create(&user)
	json.NewEncoder(w).Encode(&user)
}
