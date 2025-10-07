package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//Notice lack of newline
		fmt.Print("Pokedex > ")

		//Progress scanner to text object or loop again if missing
		if !scanner.Scan() {
			continue
		}
		//If blank input loop again
		input := scanner.Text()
		if input == "" {
			continue
		}
		cleanedInput := cleanInput(input)

		//Match command entered to cliCommand struct and handle
		// noexist
		cmd, exists := getSupportedCommands()[cleanedInput[0]]

		if exists {
			//call function in callback
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error calling %s: %v\n", cmd.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	stringLower := strings.ToLower(text)
	return strings.Fields(stringLower)
}

// Registry of CLI commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Define our supportedCommands and register
// cliCommands
func getSupportedCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return supportedCommands
}
