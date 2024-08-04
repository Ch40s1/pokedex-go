package pokeapi

// note we use a pointer to a string for next and Previous becuase
// it can sometime be null. So when it is not then it is nill
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
