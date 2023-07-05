package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xamma/gokedex/database"
	"github.com/xamma/gokedex/models"

)

func CreatePokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pokemon models.Pokemon
		if err := c.ShouldBindJSON(&pokemon); err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// Remove the current time initialization here
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		err := database.CreatePokemon(dbpool, tableName, &pokemon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pokemon created successfully"})
	}
}

func GetPokemonsHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pokemons, err := database.GetPokemons(dbpool, tableName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Pokemon records"})
			return
		}

		c.JSON(http.StatusOK, pokemons)
	}
}