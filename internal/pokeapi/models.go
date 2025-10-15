package pokeapi

// NamedAPIResource - each named API resource will have a Name and URL
// either in an array inside a NamedAPIResourceList (from calling a named API resource
// without specifying an ID or name) or nested inside a call to a named resource specifying
// an id/name for a variety of different resource - i.e this is a general/common model.
type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// NamedAPIResourceList  - this is when a named resource is called without a resource id or name.
// A paginated list (default 20 per page) is returned. The fields by default are
// count int, next string, previous string, results (list of namedAPIResource)
// NB: This used to be called LocationAreaResponse/LocationAreaPage. The new name represents the general
// structure of the objects/resources returned by the API.
type NamedAPIResourceList struct {
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

// Pokemon encounters are returned with a LocationAreaResponse. One of the fields is a
// pokemon structured as a NamedAPIResource (name, url)
type PokemonEncounter struct {
	Pokemon NamedAPIResource
}

// Calling the location-area endpoint with a name or ID returns data on a specific location-area
type LocationArea struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Pokemon_Encounters []PokemonEncounter `json:"pokemon_encounters"`
}

// Contain Pokemon Stats
type PokemonStat struct {
	Stat_info NamedAPIResource `json:"stat"`
	Base_stat int              `json:"base_stat"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

// Pokemon response when calling named endpoint with an ID - contains far more info than the
// NamedAPIResource inside PokemonEncounter
type Pokemon struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Height          int           `json:"height"`
	Weight          int           `json:"weight"`
	Stats           []PokemonStat `json:"stats"`
	Types           []PokemonType `json:"types"`
	Base_Experience int           `json:"base_experience"`
}
