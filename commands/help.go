package commands

import (
	"fmt"

	"github.com/admiralyeoj/pokedexcli/types"
)

func Help(cfg *types.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
