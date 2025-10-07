package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}

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
			//call function in callback - passing config pointer
			// note we update the values in config inside the called function
			// via config pointer.
			err := cmd.callback(config)
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

// command Config
type Config struct {
	Next     *url.URL
	Previous *url.URL
}

//func (c *Config) updateNext(url *url.URL) {
//	c.Next = url
//}

//func (c *Config) updatePrevious(url *url.URL) {
//
//	c.Previous = url
//}

// Registry of CLI commands
type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

// LocationArea - contains the result data for each location
// in the array (slice) of results returned in a LocationAreaResponse
type LocationArea struct {
	Name string
	Url  string
}

// LocationAreaResponse - api response struct for location-areas
type LocationAreaResponse struct {
	Next     string
	Previous string
	Results  []LocationArea
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
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
	}
	return supportedCommands
}
