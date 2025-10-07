package pokeapi

import (
	"net/url"
)

// command Config
type Config struct {
	Next     *url.URL
	Previous *url.URL
}
