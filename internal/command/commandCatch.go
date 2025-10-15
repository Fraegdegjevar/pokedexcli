package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandCatch(conf *pokeapi.Config, PokemonName []string) error {
	if len(PokemonName) < 1 {
		return fmt.Errorf("must supply a pokemon name")
	}

	// Catch pokemon prints to terminal, writes to pokedex. commandCatch calls from the commandline only.
	err := conf.CatchPokemon(PokemonName[0])
	if err != nil {
		return err
	}

	return nil
}
