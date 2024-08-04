package main

import (
	"time"

	"github.com/Ch40s1/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	// starts the cli
	startRepl(cfg)
}
