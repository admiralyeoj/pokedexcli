package commands

import (
	"errors"
	"fmt"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Explore(cfg *configs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	resourceResp, err := cfg.PokeApiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + name)
	fmt.Println("Found Pokemon:")

	for _, enc := range resourceResp.PokemonEncounters {
		fmt.Println("- " + enc.Pokemon.Name)
	}

	return nil
}
