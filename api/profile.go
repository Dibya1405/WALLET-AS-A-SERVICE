package api

import (
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/datastore"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user entity.User

	result := datastore.GetDbHandle().First(&user, "id=?", params["id"])
	if result.Error != nil {
		json.NewEncoder(w).Encode("<h1>You have entered invalid key</h1>")
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}
