package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandExplore(conf *pokeapi.Config, args []string) error {

	// Check input args - we need one and only one
	if len(args) != 1 {
		return fmt.Errorf("only one argument, the location-area name, should be supplied")
	}
	if args[0] == "" {
		return fmt.Errorf("blank location-area name supplied")
	}

	locationArea, err := conf.GetLocationArea(args[0])

	if err != nil {
		return err
	}

	for _, encounter := range locationArea.Pokemon_Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
