package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xamma/gokedex/models"
)

// create the Pokemontable for usage in the database initialization
// here we define the Primary key, as well as unique values.
func CreatePokemonTable(dbpool *pgxpool.Pool, tablename string) error {
    _, err := dbpool.Exec(
        context.Background(),
        fmt.Sprintf(`
            CREATE TABLE IF NOT EXISTS %s (
                id SERIAL PRIMARY KEY,
                name TEXT UNIQUE,
                type TEXT,
                caught_date TIMESTAMP
            )
        `, tablename),
    )
    if err != nil {
        return err
    }

    return nil
}

// CreatePokemon creates a new Pokémon record in the database
func CreatePokemon(dbpool *pgxpool.Pool, tablename string, pokemon *models.Pokemon) error {
	pokemon.CaughtDate = time.Now().UTC() // Initialize CaughtDate with current time

    _, err := dbpool.Exec(
        context.Background(),
        fmt.Sprintf("INSERT INTO %s (name, type, caught_date) VALUES ($1, $2, $3)", tablename),
        pokemon.Name,
        pokemon.Type,
        pokemon.CaughtDate,
    )
    if err != nil {
        return err
    }

    return nil
}
// GetPokemon retrieves a Pokémon record from the database based on the provided ID
func GetPokemon(dbpool *pgxpool.Pool, tablename string, name string) (*models.Pokemon, error) {
	var pokemon models.Pokemon

    err := dbpool.QueryRow(
        context.Background(),
        fmt.Sprintf("SELECT name, type, caught_date FROM %s WHERE name = $1", tablename),
        name,
    ).Scan(&pokemon.Name, &pokemon.Type, &pokemon.CaughtDate)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

// GetPokemons returns a slice of Pokémon record from the database
func GetPokemons(dbpool *pgxpool.Pool, tablename string) ([]models.Pokemon, error) {
    rows, err := dbpool.Query(
        context.Background(),
        fmt.Sprintf("SELECT name, type, caught_date FROM %s", tablename),
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pokemons []models.Pokemon
    for rows.Next() {
        var pokemon models.Pokemon
        if err := rows.Scan(&pokemon.Name, &pokemon.Type, &pokemon.CaughtDate); err != nil {
            return nil, err
        }
        pokemons = append(pokemons, pokemon)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return pokemons, nil
}

// UpdatePokemon updates an existing Pokémon record in the database
// this uses the name to update, this is okay since its unique in the database.
func UpdatePokemon(dbpool *pgxpool.Pool, tablename string, pokemon *models.Pokemon) error {
    _, err := dbpool.Exec(
        context.Background(),
        fmt.Sprintf("UPDATE %s SET type = $1, caught_date = $2 WHERE name = $3", tablename),
        pokemon.Type,
        pokemon.CaughtDate,
        pokemon.Name,
    )
    if err != nil {
        return err
    }

    return nil
}

// UpdatePokemonbyID updates based on the ID in the database.
func UpdatePokemonbyID(dbpool *pgxpool.Pool, tablename string, id int, pokemon *models.Pokemon) error {
    _, err := dbpool.Exec(
        context.Background(),
        fmt.Sprintf("UPDATE %s SET name = $1, type = $2, caught_date = $3 WHERE id = $4", tablename),
        pokemon.Name,
        pokemon.Type,
        pokemon.CaughtDate,
        id,
    )
    if err != nil {
        return err
    }

    return nil
}

// DeletePokemon deletes a Pokémon record from the database based on the provided name
// The Exec function is used, because there are no rows returned
func DeletePokemon(dbpool *pgxpool.Pool, tablename string, name string) error {
    _, err := dbpool.Exec(
		context.Background(),
		fmt.Sprintf("DELETE FROM %s WHERE name = $1", tablename),
		name,
	)

	if err != nil {
		return err
	}

	return nil
}
