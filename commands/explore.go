package commands

import (
	"fmt"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Explore(cfg *configs.Config, name string) error {
	// name = "mt-coronet-1f-route-207"
	resourceResp, err := cfg.PokeApiClient.ResourceList(name)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + name)
	fmt.Println("Found Pokemon:")

	for _, row := range resourceResp.PokemonEncounters {
		fmt.Println("- " + row.Pokemon.Name)
	}

	return nil
}
