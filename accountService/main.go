package main

import (
	"fmt"

	"github.com/damilarelana/goSystemDesign/accountService/dbClients"
	middleware "github.com/damilarelana/goSystemDesign/middleware"
)

var appName = "accountService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	middleware.InitializeWebServer("8080")
}

// initializeBoltClient() using the boltDb interface
// - creates a boltDB instance
// - opens a OpenBoltDB connection
// - seeds the database with a few accounts
//
func initializeBoltClient() {
	middleware.DBClient = &dbClients.BoltClient{} // create boltDB instance by passing an implementation of the interface to it
	middleware.DBClient.OpenBoltDB()
	middleware.DBClient.Seed()
}
