package command

import (
	"fmt"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandMap(conf *pokeapi.Config) error {
	//Default behaviour is to return batches of 20 location-areas.
	//Use the next URL stored in conf if it exists and update next/previous
	// Else default to the base URL and update next

	fullURL := conf.Next

	locationAreas, err := conf.GetLocationAreas(fullURL)
	if err != nil {
		return err
	}

	//Now fetch the names of the location-areas within results slice
	for _, locarea := range locationAreas.Results {
		fmt.Println(locarea.Name)
	}
	return nil
}
