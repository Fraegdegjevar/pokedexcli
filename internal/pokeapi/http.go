package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func GetLocationAreas(fullURL *url.URL, conf *Config) (LocationAreaResponse, error) {
	//Build request
	if fullURL == nil {
		fullURL, _ = url.Parse(baseURL + "/location-area/")
	}
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		fmt.Println("Error generating NewRequest")
		return LocationAreaResponse{}, err
	}

	//initialise HTTP client
	client := &http.Client{}
	//Do request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching map info from pokeapi")
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	//Decode response
	var data LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding response body")
		return LocationAreaResponse{}, err
	}
	return data, nil
}
