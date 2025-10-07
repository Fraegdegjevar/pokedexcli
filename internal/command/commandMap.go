package command

import (
	"fmt"
	"net/url"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func commandMap(conf *pokeapi.Config) error {
	//Default behaviour is to return batches of 20 location-areas.
	fullURL, err := url.Parse("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("Error parsing base url!")
		return err
	}
	//Use the next URL stored in conf if it exists and update next/previous
	// Else default to the base URL and update next
	if conf.Next != nil {
		fullURL = conf.Next
	}

	locationAreas, err := pokeapi.GetLocationAreas(fullURL, conf)
	if err != nil {
		return err
	}

	//Store the next and previous urls in the API response in conf
	conf.Next, err = url.Parse(locationAreas.Next)
	if err != nil {
		fmt.Println("Error parsing response next URL field as a url.URL")
		return err
	}
	conf.Previous, err = url.Parse(locationAreas.Previous)
	if err != nil {
		fmt.Println("Error parsing response previous URL field as a url.URL")
		return err
	}

	//Now fetch the names of the location-areas within results slice
	for _, locarea := range locationAreas.Results {
		fmt.Println(locarea.Name)
	}
	return nil
}
