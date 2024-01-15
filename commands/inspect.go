package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Inspect(cfg *configs.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide the name of the pokemon you would like to inspect")
	}

	name := args[0]
	pokemon, ok := cfg.CaughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: " + pokemon.Name)
	fmt.Println("Height: " + strconv.Itoa(pokemon.Height))
	fmt.Println("Weight: " + strconv.Itoa(pokemon.Weight))

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
