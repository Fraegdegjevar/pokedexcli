package command

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	//Build request
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		fmt.Println("Error generating NewRequest")
		return err
	}

	//initialise HTTP client
	client := &http.Client{}
	//Do request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching map info from pokeapi")
		return err
	}

	//Decode response
	var data pokeapi.LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding response body")
		return err
	}
	//Update config with the new next and previous pages from response
	conf.Next, err = url.Parse(data.Next)
	if err != nil {
		fmt.Println("Error parsing response Next URL field as a url.URL")
		return err
	}
	conf.Previous, err = url.Parse(data.Previous)
	if err != nil {
		fmt.Println("Error parsing response Previous URL field as a url.URL")
		return err
	}

	//Print the location area names
	for _, locarea := range data.Results {
		fmt.Println(locarea.Name)
	}

	return nil
}
