package main

import (
	"WALLET-AS-A-SERVICE/datastore"
	"WALLET-AS-A-SERVICE/server"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the project : WALLET-AS-A-SERVICE")
	datastore.InitializeDb()
	server.StartServer()
	fmt.Println("Listening port at 4000...")
}
