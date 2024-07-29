package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to match the structure of the JSON response
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

func getPokeApi() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatalf("Error getting request from API: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// Decode the JSON response into the PokeApiResponse struct
	var result PokeApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Print just the names of the locations
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}
