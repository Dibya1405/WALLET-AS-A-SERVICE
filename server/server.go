package server

import (
	"WALLET-AS-A-SERVICE/logger"
	"log"
	"net/http"

	"WALLET-AS-A-SERVICE/router"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	router.InitializeRouter(r)
	logger.GetMyLogger().Info("Routers are initialised")
	logger.GetMyLogger().Warn("Server started successfully")
	log.Fatal(http.ListenAndServe(":4000", r))
}
