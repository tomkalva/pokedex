package main

import (
	"time"

	"pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		previousURL:   nil,
		nextURL:       nil,
		pokedex:       map[string]pokeapi.PokemonResponse{},
	}

	startRepl(cfg)
}
