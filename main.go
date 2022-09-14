package main

import (
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/logger"
	"WALLET-AS-A-SERVICE/server"
)

func main() {
	// If you want to log in a particular file ---

	//file, err := os.Create("./file.log")
	//utils.CheckError(err)
	//defer file.Close()
	//log.SetOutput(file)

	logger.GetMyLogger().Info("Welcome to the project : WALLET-AS-A-SERVICE")
	datastore.InitializeDb()
	server.StartServer()
	logger.GetMyLogger().Warn("Listening port at 4000...")
}
