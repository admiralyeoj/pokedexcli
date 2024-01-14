package commands

import (
	"errors"
	"fmt"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Mapb(cfg *configs.Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeApiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
