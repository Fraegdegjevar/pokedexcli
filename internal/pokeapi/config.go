package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokecache"
)

// command Config
type Config struct {
	Next                 *url.URL
	Previous             *url.URL
	GetLocationAreasFunc func(*url.URL) (LocationAreaResponse, error)
	Cache                *pokecache.Cache
}

func (c *Config) UpdatePagination(resp *LocationAreaResponse) error {
	var err error
	c.Next, err = url.Parse(resp.Next)
	if err != nil {
		fmt.Println("Error parsing response next URL field as a url.URL")
		return err
	}
	c.Previous, err = url.Parse(resp.Previous)
	if err != nil {
		fmt.Println("Error parsing response previous URL field as a url.URL")
		return err
	}

	return nil
}

func (c *Config) GetLocationAreas(u *url.URL) (LocationAreaResponse, error) {
	var resp LocationAreaResponse

	// Guard null url value
	if u == nil {
		u, _ = url.Parse(baseURL + "/location-area/?offset=0&limit=20")
	}

	//Test if result in cache and parse if necessary
	page, found := c.Cache.Get(u.String())
	if found {
		fmt.Printf("Cache hit on url: %v\n", u)
		err := json.Unmarshal(page, &resp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		err = c.UpdatePagination(&resp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return resp, nil
	}
	fmt.Printf("Cache miss on url: %v\n", u)

	resp, err := RequestLocationAreas(u)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	page, err = json.Marshal(&resp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.Cache.Add(u.String(), page)
	c.UpdatePagination(&resp)

	return resp, err
}
