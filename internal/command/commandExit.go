package command

import (
	"fmt"
	"os"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

// commandfunctions
func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
