package commands

import (
	"fmt"

	"github.com/admiralyeoj/pokedexcli/types"
)

func Map(cfg *types.Config) error {
	locationsResp, err := cfg.PokeApiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
