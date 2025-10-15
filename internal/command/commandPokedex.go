package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandPokedex(conf *pokeapi.Config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for key := range conf.Pokedex {
		fmt.Println("  -", key)
	}

	return nil
}
