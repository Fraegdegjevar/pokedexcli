package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandInspect(conf *pokeapi.Config, PokemonName []string) error {
	if len(PokemonName) < 1 {
		return fmt.Errorf("you must supply a pokemon to inspect")
	}

	pokemon, err := conf.InspectPokemon(PokemonName[0])

	//if error in finding pokemon, or if it does not exist in the pokedex
	if err != nil {
		return err
	}

	// Print the fields we care about
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	// Loop through stats
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", s.Stat_info.Name, s.Base_stat)
	}
	// Loop through Types
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}
	return nil
}
