package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xamma/gokedex/database"
	"github.com/xamma/gokedex/models"
)

// CreatePokemonHandler             godoc
// @Summary      Add Pokemon
// @Description  Adds a Pokemon to the Pokedex
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Param        pokemon body models.Pokemon true "Pokemon object"
// @Success      200 {string} string "Pokemon created successfully"
// @Router       /pokemon [post]
func CreatePokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pokemon models.Pokemon
		if err := c.ShouldBindJSON(&pokemon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := database.CreatePokemon(dbpool, tableName, &pokemon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pokemon created successfully"})
	}
}

// GetPokemonsHandler             godoc
// @Summary      Get Pokemons.
// @Description  Gets all pokemons from the pokedex.
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Success      200 {string} string
// @Router       /pokemons [get]
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

// GetPokemonHandler             godoc
// @Summary      Get Pokemon
// @Description  Gets a Pokemon from the Pokedex by name
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Param        name path string true "Pokemon name"
// @Success      200 {object} models.Pokemon
// @Failure      404 {string} string "Pokemon not found"
// @Failure      500 {string} string "Failed to retrieve Pokemon record"
// @Router       /pokemon/{name} [get]
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

// DeletePokemonHandler             godoc
// @Summary      Deletes Pokemon
// @Description  Deletes a Pokemon from the Pokedex by name
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Param        name path string true "Pokemon name"
// @Success      200 {string} string "Successfully deleted Pokemon"
// @Failure      500 {string} string "Failed to delete Pokemon"
// @Router       /pokemon/{name} [delete]
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

// PutNamePokemonHandler             godoc
// @Summary      Update Pokemon by name
// @Description  Updates a Pokemon in the Pokedex by name
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Param        pokemon body models.Pokemon true "Pokemon object"
// @Success      200 {string} string "Successfully updated Pokemon"
// @Failure      400 {object} string "Failed to update Pokemon"
// @Failure      500 {string} string "Failed to update Pokemon"
// @Router       /pokemon [put]
func PutNamePokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pokemon models.Pokemon

		if err := c.ShouldBindJSON(&pokemon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := database.UpdatePokemon(dbpool, tableName, &pokemon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Pokemon!"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully updated Pokemon."})
		}
	}
}

// @Summary      Update Pokemon by ID
// @Description  Updates a Pokemon in the Pokedex by ID
// @Tags         pokedex
// @Accept       json
// @Produce      json
// @Param        id path string true "Pokemon ID"
// @Param        pokemon body models.Pokemon true "Pokemon object"
// @Success      200 {string} string "Successfully updated Pokemon"
// @Failure      400 {object} string "Failed to update Pokemon"
// @Failure      500 {string} string "Failed to update Pokemon"
// @Router       /pokemon/{id} [put]
func PutIDPokemonHandler(dbpool *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Convert the ID to an integer
		pokemonID, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var pokemon models.Pokemon

		if err := c.ShouldBindJSON(&pokemon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = database.UpdatePokemonbyID(dbpool, tableName, pokemonID, &pokemon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Pokemon!"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully updated Pokemon."})
		}
	}
}