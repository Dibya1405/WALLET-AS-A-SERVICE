package server

import (
	"fmt"
	"log"
	"net/http"

	"WALLET-AS-A-SERVICE/router"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	router.InitializeRouter(r)
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started successfully")
}
