package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xamma/gokedex/database"
)

func main() {

	// Create a DB connection pool
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbpool, err := database.createConnPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to the database instance: %v", err)
	}
	defer dbpool.Close()

	// Check if the database already exists, othwise create it
	var dbname string = "poke"
	err = createDatabaseIfNotExists(dbpool, dbname)
	if err != nil {
		log.Fatalf("Unable to create database: %v", err)
	}

	// finally connect to the database
	conn, err := connectToDatabase(os.Getenv("DATABASE_URL"), dbname)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close()
}

func createConnPool(databaseURL string) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	return connPool, nil
}

func createDatabaseIfNotExists(dbpool *pgxpool.Pool, dbname string) error {
	var exists bool
	err := dbpool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = dbpool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			return err
		}
	}

	return nil
}

func connectToDatabase(databaseURL string, databaseName string) (*pgxpool.Pool, error) {
	var connString string = databaseURL + databaseName
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}