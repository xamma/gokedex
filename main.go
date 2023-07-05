package main

import (
	"fmt"
	"log"

	"github.com/xamma/gokedex/config"
	"github.com/xamma/gokedex/database"
	"github.com/xamma/gokedex/handlers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/xamma/gokedex/docs"
)

// @title           Gokedex
// @version         1.0
// @description     Gokedex RestAPI with PostgreSQL.
// @termsOfService  placeholder

// @contact.name   Max
// @contact.url    xamma.github.io
// @contact.email  xam93@live.de

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1
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

	// disable debug mode
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	router := gin.Default()

	//router group
	v1 := router.Group("/api/v1")

	// Define routes
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.POST("/pokemon", handlers.CreatePokemonHandler(conn, config.DatabaseTableName))
	v1.GET("/pokemons", handlers.GetPokemonsHandler(conn, config.DatabaseTableName))
	v1.GET("/pokemon/:name", handlers.GetPokemonHandler(conn, config.DatabaseTableName))
	v1.DELETE("/pokemon/:name", handlers.DeletePokemonHandler(conn, config.DatabaseTableName))
	v1.PUT("/pokemon", handlers.PutNamePokemonHandler(conn, config.DatabaseTableName))
	v1.PUT("/pokemon/:id", handlers.PutIDPokemonHandler(conn, config.DatabaseTableName))

	port := fmt.Sprintf(":%s", config.ApiPort)
	router.Run(port)
}