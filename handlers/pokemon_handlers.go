package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

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

func GetPokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if we want to work with DB id
		// id := c.Param("id")
		pokeName := c.Param("name")

		// Convert the ID to an integer
		// pokemonID, err := strconv.Atoi(id)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		// 	return
		// }

		pokemon, err := database.GetPokemon(dbpool, tableName, pokeName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Pokemon record"})
			return
		}

		if pokemon == nil {
			// Handle the case when no Pokemon is found with the name
			c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

func DeletePokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pokeName := c.Param("name")

		err := database.DeletePokemon(dbpool, tableName, pokeName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Pokemon!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message":  fmt.Sprintf("Successfully deleted Pokemon %s. It was a great Pokemon!", pokeName)})
	}
}