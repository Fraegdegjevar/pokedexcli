package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandMapb(conf *pokeapi.Config) error {
	fullURL := conf.Previous

	//in case we are already on first page (no previous)
	if fullURL == nil || fullURL.Path == "" {
		fmt.Println("you're on the first page.")
		return nil
	}

	locationAreas, err := conf.GetLocationAreas(fullURL)
	if err != nil {
		return err
	}

	//Print the location area names
	for _, locarea := range locationAreas.Results {
		fmt.Println(locarea.Name)
	}
	return nil
}
