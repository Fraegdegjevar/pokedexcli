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
	GetLocationAreasFunc func(*url.URL) (NamedAPIResourceList, error)
	Cache                *pokecache.Cache
}

func (c *Config) UpdatePagination(resp *NamedAPIResourceList) error {
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

func (c *Config) GetLocationAreas(u *url.URL) (NamedAPIResourceList, error) {
	var resp NamedAPIResourceList
	var err error
	// Guard null url value
	if u == nil {
		u, err = url.Parse(baseURL + "/location-area/?offset=0&limit=20")
	}

	if err != nil {
		return NamedAPIResourceList{}, fmt.Errorf("error parsing URL for GetLocationAreas: %v", err)
	}

	//Test if result in cache and parse if necessary
	page, found := c.Cache.Get(u.String())
	if found {
		fmt.Printf("Cache hit on url: %v\n", u)
		err := json.Unmarshal(page, &resp)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		err = c.UpdatePagination(&resp)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		return resp, nil
	}
	fmt.Printf("Cache miss on url: %v\n", u)

	resp, err = RequestLocationAreas(u)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	page, err = json.Marshal(&resp)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	c.Cache.Add(u.String(), page)
	c.UpdatePagination(&resp)

	return resp, nil
}

// Gets a specific location area resource. To add: cache checks and caching.
func (c *Config) GetLocationArea(LocationAreaName string) (LocationArea, error) {
	// Handle empty strings
	u, err := url.Parse(baseURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error parsing url in GetLocationArea: %v", err)

	}
	// Append to url path as needed to hit correct resource
	u = u.JoinPath(LocationAreaEndpoint, LocationAreaName)

	resp, err := RequestLocationArea(u)

	if err != nil {
		return LocationArea{}, err
	}

	//NOTE: No check against cache or adding result to cache yet.
	return resp, nil
}
