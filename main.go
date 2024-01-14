package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// cliName is the name used in the repl prompts
const cliName string = "Pokedex"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// String returns a string representation of the cliCommand
func (c cliCommand) String() string {
	indent := strings.Repeat("\t", 1)
	return fmt.Sprintf("%s%s - %s", c.name, indent, c.description)
}

// Commands available to be ran
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"clear": {
			name:        "clear",
			description: "Clear the terminal",
			callback:    commandClear,
		},
		"exit": {
			name:        "exit",
			description: "Exit the " + cliName,
			callback:    commandExit,
		},
	}
}

// commandHelp informs the user about our functions
func commandHelp() error {
	commands := getCommands()
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		cliName,
	)
	for _, command := range commands {
		fmt.Println(command.String())
		// fmt.Printf("%s - %s", command.name, command.description)
	}
	// fmt.Println("help    - Show available commands")
	// fmt.Println("clear   - Clear the terminal screen")
	// fmt.Println("exit    - Closes your connection to ", cliName)

	return nil
}

// commandClear clears the terminal screen
func commandClear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}

// commandExit exits the application
func commandExit() error {
	// Implement the exit command logic here (optional)
	fmt.Println("Exiting the program...")
	os.Exit(0)
	return nil
}

// printUnkown informs the user about invalid commands
func commandUnknown(text string) {
	fmt.Println(text, ": command not found")
}

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
	defer commandUnknown(text)
}

// handleCmd parses the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	commands := getCommands()
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			// Call the callback function if it exists
			if command.callback != nil {
				err := command.callback()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}
		} else {
			// Pass the command to the parser
			handleCmd(text)
		}
		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
