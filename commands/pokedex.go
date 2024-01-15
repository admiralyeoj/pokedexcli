package commands

import (
	"errors"
	"fmt"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Pokedex(cfg *configs.Config) error {
	caught := cfg.CaughtPokemon

	if len(caught) == 0 {
		return errors.New("No pokemon have been caught.")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range caught {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
