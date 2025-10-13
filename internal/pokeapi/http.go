package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL              = "https://pokeapi.co/api/v2"
	LocationAreaEndpoint = "/location-area/"
)

func RequestLocationAreas(fullURL *url.URL) (NamedAPIResourceList, error) {
	//Build request
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		fmt.Println("Error generating NewRequest")
		return NamedAPIResourceList{}, err
	}

	//initialise HTTP client
	client := &http.Client{
		// Set a timeout for receiving a response that accounts for network latency on API side.
		Timeout: 10 * time.Second,
	}
	//Do request - note that err only returns non nil
	// if there was an error wit hthe http exchange. So if we receive a response
	// even if it is an error code, err is nil. We need to explicitly handle error codes.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching map info from pokeapi: %v\n", err)
		return NamedAPIResourceList{}, err
	}
	defer resp.Body.Close()

	// Test for application errors i.e http error codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return NamedAPIResourceList{}, fmt.Errorf("unexpected HTTP status: %v", resp.StatusCode)
	}

	//Decode response
	var data NamedAPIResourceList
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding response body")
		return NamedAPIResourceList{}, err
	}

	return data, nil
}

func RequestLocationArea(fullURL *url.URL) (LocationArea, error) {
	// Build request
	req, err := http.NewRequest("GET", fullURL.String(), nil)

	if err != nil {
		return LocationArea{}, fmt.Errorf("error generating request: %v", err)
	}

	//Initialise HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//Perform request
	resp, err := client.Do(req)

	if err != nil {
		return LocationArea{}, fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Test for application errors i.e http error codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return LocationArea{}, fmt.Errorf("unexpected HTTP status: %v", resp.StatusCode)
	}

	// Decode
	var locationArea LocationArea
	err = json.NewDecoder(resp.Body).Decode(&locationArea)

	if err != nil {
		return LocationArea{}, fmt.Errorf("error decoding response: %v", err)
	}

	return locationArea, nil
}
