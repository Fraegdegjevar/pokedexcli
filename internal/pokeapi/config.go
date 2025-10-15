package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokecache"
)

const (
	baseURL              = "https://pokeapi.co/api/v2"
	LocationAreaEndpoint = "/location-area/"
	PokemonEndpoint      = "/pokemon/"
)

// command Config
type Config struct {
	Next                 *url.URL
	Previous             *url.URL
	GetLocationAreasFunc func(*url.URL) (NamedAPIResourceList, error)
	Cache                *pokecache.Cache
	Pokedex              map[string]Pokemon
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

// Get pokemon from API or cache
func (c *Config) CatchPokemon(PokemonName string) error {
	// construct url
	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("could not parse baseURL: %v", err)
	}
	u = u.JoinPath(PokemonEndpoint, PokemonName)
	// Request pokemon
	pokemon, err := RequestPokemon(u)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	//Check if pokemon already caught! We still catch but this is helpful.
	_, caught := c.Pokedex[pokemon.Name]
	fmt.Printf("%v already caught: %v\n", pokemon.Name, caught)

	// Determine probability of catching pokemon: 100 - base_exp/5 %
	// We roll a 100 sided die. If the roll is < prob, success, else fail.
	// As the die returns 0 through 99, need to have succes on < not <= prob
	// I.e a prob of 63 means we need 63/100 to be success. 0-62 succeed, 63-99 fail
	prob := 100 - (pokemon.Base_Experience / 5)
	roll := rand.Intn(100)

	if roll < prob {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		// add pokemon to pokemon map here...
		c.Pokedex[pokemon.Name] = pokemon
		return nil
	}
	// If pokemon escaped...
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}

// check if in pokedex. Retrieve if so else propagate error to calling function commandInspect
func (c *Config) InspectPokemon(PokemonName string) (Pokemon, error) {
	// input check
	if len(PokemonName) < 1 {
		return Pokemon{}, fmt.Errorf("you must supply a pokemon to inspect")
	}

	//in pokedex?
	pokemon, found := c.Pokedex[PokemonName]

	if !found {
		return Pokemon{}, fmt.Errorf("you have not caught that Pokemon")
	}

	return pokemon, nil
}
