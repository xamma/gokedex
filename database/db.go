package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnPool(databaseURL string) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	return connPool, nil
}

func CreateDatabaseIfNotExists(dbpool *pgxpool.Pool, dbname string) error {
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

func ConnectToDatabase(databaseURL string, databaseName string) (*pgxpool.Pool, error) {
	var connString string = databaseURL + databaseName
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}