package main

import (
	"os"
	"log"

	"github.com/xamma/gokedex/database"
)

func main() {

	// Create a DB connection pool
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbpool, err := database.CreateConnPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to the database instance: %v", err)
	}
	defer dbpool.Close()

	// Check if the database already exists, othwise create it
	var dbname string = "poke"
	err = database.CreateDatabaseIfNotExists(dbpool, dbname)
	if err != nil {
		log.Fatalf("Unable to create database: %v", err)
	}

	// finally connect to the database
	conn, err := database.ConnectToDatabase(os.Getenv("DATABASE_URL"), dbname)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close()
}
