package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

// Registry of CLI commands
type cliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config) error
}

// Define our supportedCommands and register
// cliCommands
func GetSupportedCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of the next 20 location areas in the Pokemon world.",
			// Closure to allow us to return a function of more than just *pokeapi.Config
			Callback: commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			Callback:    commandMapb,
		},
	}
	return supportedCommands
}

// Match input (first word) to supported commands and callback.
func ExecuteCommand(input []string, config *pokeapi.Config) error {
	//Match command entered to cliCommand struct and handle
	// noexist
	cmd, exists := GetSupportedCommands()[input[0]]

	if exists {
		//call function in callback - passing config pointer
		// note we update the values in config inside the called function
		// via config pointer.
		err := cmd.Callback(config)
		if err != nil {
			return fmt.Errorf("error calling %s: %v", cmd.Name, err)
		}
	} else {
		fmt.Println("Unknown command")
	}
	return nil
}
