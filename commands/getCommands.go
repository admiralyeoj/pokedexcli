package commands

import (
	"github.com/admiralyeoj/pokedexcli/configs"
)

type cliCommand struct {
	name         string
	description  string
	Callback     func(cfg *configs.Config) error
	CallbackArgs func(cfg *configs.Config, args ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			Callback:    Map,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations. This does not work on the first page.",
			Callback:    Mapb,
		},
		"explore": {
			name:         "explore",
			description:  "Enter a location name and get the pokemon that can be encountered.",
			CallbackArgs: Explore,
		},
		"catch": {
			name:         "catch",
			description:  "Have a chance to catch the pokemon of your choosing.",
			CallbackArgs: Catch,
		},
		"inspect": {
			name:         "inspect",
			description:  "Look at pokemon information of the pokemon that you have caught.",
			CallbackArgs: Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Output of all the pokemon that you have caught.",
			Callback:    Pokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"clear": {
			name:        "clear",
			description: "Clear the terminal",
			Callback:    Clear,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
	}
}
