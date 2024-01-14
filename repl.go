package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/admiralyeoj/pokedexcli/commands"
	"github.com/admiralyeoj/pokedexcli/types"
)

// cliName is the name used in the repl prompts
const cliName string = "Pokedex"

func startRepl(cfg *types.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(cliName + " > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println(commandName, ": command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
