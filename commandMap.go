package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func commandMap(conf *Config) error {
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

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		fmt.Println("Error generating NewRequest")
		return err
	}

	//initialise http client
	client := &http.Client{}
	//initialise response struct
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching map info from pokeapi")
		return err
	}
	defer resp.Body.Close()

	//decode response
	var data LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding response body")
		return err
	}

	//Store the next and previous urls in the API response in conf
	conf.Next, err = url.Parse(data.Next)
	if err != nil {
		fmt.Println("Error parsing response next URL field as a url.URL")
		return err
	}
	conf.Previous, err = url.Parse(data.Previous)
	if err != nil {
		fmt.Println("Error parsing response previous URL field as a url.URL")
		return err
	}

	//Now fetch the names of the location-areas within results slice
	for _, locarea := range data.Results {
		fmt.Println(locarea.Name)
	}

	return nil
}
