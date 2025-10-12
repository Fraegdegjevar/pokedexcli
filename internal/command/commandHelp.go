package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config) error {
	// Print welcome and  usage instructions for our supportedCommands
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	// Get map of supportedcommands with function
	for _, cmd := range GetSupportedCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
