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
			return NamedAPIResourceList{}, err
		}
		err = c.UpdatePagination(&resp)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		return resp, nil
	}
	fmt.Printf("Cache miss on url: %v\n", u)

	resp, err := RequestLocationAreas(u)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	page, err = json.Marshal(&resp)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	c.Cache.Add(u.String(), page)
	c.UpdatePagination(&resp)

	return resp, err
}
