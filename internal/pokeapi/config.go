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
	var page NamedAPIResourceList
	var err error
	// Guard null url value
	if u == nil {
		u, err = url.Parse(baseURL + "/location-area/?offset=0&limit=20")
	}

	if err != nil {
		return NamedAPIResourceList{}, fmt.Errorf("error parsing URL for GetLocationAreas: %v", err)
	}

	//Test if result in cache and parse if necessary
	resp, exists := c.Cache.Get(u.String())
	if exists {
		fmt.Printf("Cache hit on url: %v\n", u)
		err := json.Unmarshal(resp, &page)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		err = c.UpdatePagination(&page)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		return page, nil
	}
	fmt.Printf("Cache miss on url: %v\n", u)

	page, err = requestLocationAreas(u)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	resp, err = json.Marshal(&page)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	c.Cache.Add(u.String(), resp)
	c.UpdatePagination(&page)

	return page, nil
}

// Gets a specific location area resource. To add: cache checks and caching.
func (c *Config) GetLocationArea(LocationAreaName string) (LocationArea, error) {
	var locationArea LocationArea

	// Handle empty strings
	u, err := url.Parse(baseURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error parsing url in GetLocationArea: %v", err)

	}
	// Append to url path as needed to hit correct resource
	u = u.JoinPath(LocationAreaEndpoint, LocationAreaName)
	// Check for presence of URL as key in cache
	resp, exists := c.Cache.Get(u.String())
	if exists {
		fmt.Printf("Cache hit on url: %v\n", u)
		err = json.Unmarshal(resp, &locationArea)
		if err != nil {
			return LocationArea{}, nil
		}
		// No pagination update required.
		return locationArea, nil
	}

	// If not in cache, call API
	fmt.Printf("Cache miss on url: %v\n", u)
	locationArea, err = requestLocationArea(u)
	if err != nil {
		return LocationArea{}, nil
	}

	//Add to cache - marshal
	resp, err = json.Marshal(locationArea)
	if err != nil {
		return LocationArea{}, nil
	}
	c.Cache.Add(u.String(), resp)

	return locationArea, nil
}
