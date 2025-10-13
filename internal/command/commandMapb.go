package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandMapb(conf *pokeapi.Config, _ []string) error {
	//in case we are already on first page (no previous)
	if conf.Previous == nil || conf.Previous.Path == "" {
		fmt.Println("you're on the first page.")
		return nil
	}

	locationAreas, err := conf.GetLocationAreas(conf.Previous)
	if err != nil {
		return err
	}

	//Print the location area names
	for _, locarea := range locationAreas.Results {
		fmt.Println(locarea.Name)
	}
	return nil
}
