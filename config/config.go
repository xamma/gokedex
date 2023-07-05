package config

import (
	"os"

	"github.com/xamma/gokedex/models"
)

func LoadConfig() (*models.AppConfig, error) {
	defaultEnvVars := map[string]string{
		"DATABASE_URL": "postgres://postgres:mysecretpassword@localhost:5432/",
		"DATABASE_NAME": "pokedex",
		"DATABASE_TABLE_NAME": "pokemon",
		"API_PORT": "8080",
	}

	databaseUrl := getEnv("DATABASE_URL", defaultEnvVars)
	databaseName := getEnv("DATABASE_NAME", defaultEnvVars)
	databaseTableName := getEnv("DATABASE_TABLE_NAME", defaultEnvVars)
	apiPort := getEnv("API_PORT", defaultEnvVars)

	config := &models.AppConfig{
		DatabaseUrl: databaseUrl,
		DatabaseName: databaseName,
		DatabaseTableName: databaseTableName,
		ApiPort: apiPort,
	}

	return config, nil
}

func getEnv(key string, defaultValues map[string]string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValues[key]
	}
	return value
}