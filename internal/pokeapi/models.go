package pokeapi

// LocationArea - contains the result data for each location
// in the array (slice) of results returned in a LocationAreaResponse
type LocationArea struct {
	Name string
	Url  string
}

// LocationAreaResponse - api response struct for location-areas
type LocationAreaResponse struct {
	Next     string
	Previous string
	Results  []LocationArea
}
