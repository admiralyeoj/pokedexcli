package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Catch(cfg *configs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide the name of the pokemon you would like to catch")
	}

	name := args[0]
	pokemon, err := cfg.PokeApiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	name = pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	// Calculate the catch rate
	catchRate := calculateCatchRate(pokemon.BaseExperience)

	// fmt.Println(, pokemon.BaseExperience)

	// Simulate a catch attempt
	if attemptCatch(catchRate) {
		fmt.Printf("%s was caught!", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!", pokemon.Name)
	}
	cfg.CaughtPokemon[pokemon.Name] = pokemon

	fmt.Println()
	return nil
}

func calculateCatchRate(baseExperience int) float64 {
	// Define your own formula or logic here
	// For example, a simple formula could be:
	// Catch Rate = 1 / (1 + baseExperience/1000)
	return 1.0 / (1.0 + float64(baseExperience)/1000.0)
}

func attemptCatch(catchRate float64) bool {
	// Generate a random number between 0 and 1
	randNum := rand.Float64()

	// Check if the random number is less than the catch rate
	return randNum < catchRate
}
