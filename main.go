package main

import (
	"log"

	"github.com/xamma/gokedex/database"
	"github.com/xamma/gokedex/config"
)

func main() {

	// load configs and overwrite defaults with ENVs
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize database connection
	conn, err := database.InitDatabase(config)
	if err != nil {
		log.Fatalf("Unable to initialize the database: %v", err)
	}
	defer conn.Close()
}
