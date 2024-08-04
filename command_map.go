package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

var offset int
var limit = 20
var mu sync.Mutex

func getPokeApi(offset int) ([]Location, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location/?offset=%d&limit=%d", offset, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting request from API: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	var result PokeApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return result.Results, nil
}

func getMapNext() error {
	mu.Lock()
	defer mu.Unlock()

	locations, err := getPokeApi(offset)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	offset += limit
	return nil
}

func getMapPrevious() error {
	mu.Lock()
	defer mu.Unlock()

	if offset == 0 {
		return fmt.Errorf("cannot call mapb: In first page")
	}

	offset -= limit
	locations, err := getPokeApi(offset)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
