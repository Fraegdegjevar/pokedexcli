package main

import (
	"fmt"
)

func commandHelp(config *Config) error {
	// Print welcome and  usage instructions for our supportedCommands
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	// Get map of supportedcommands with function
	for _, cmd := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
