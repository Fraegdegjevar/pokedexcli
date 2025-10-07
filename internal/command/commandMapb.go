package command

import (
	"fmt"
	"net/url"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandMapb(conf *pokeapi.Config) error {
	var fullURL *url.URL
	if conf.Previous != nil && conf.Previous.Path != "" {
		fullURL = conf.Previous
	} else {
		//in case we are already on first page (no previous)
		fmt.Println("you're on the first page.")
		return nil
	}

	locationAreas, err := pokeapi.GetLocationAreas(fullURL, conf)

	//Update config with the new next and previous pages from response
	conf.Next, err = url.Parse(locationAreas.Next)
	if err != nil {
		fmt.Println("Error parsing response Next URL field as a url.URL")
		return err
	}
	conf.Previous, err = url.Parse(locationAreas.Previous)
	if err != nil {
		fmt.Println("Error parsing response Previous URL field as a url.URL")
		return err
	}

	//Print the location area names
	for _, locarea := range locationAreas.Results {
		fmt.Println(locarea.Name)
	}
	return nil
}
