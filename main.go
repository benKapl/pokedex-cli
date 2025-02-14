package main

import (
	"time"

	"github.com/benKapl/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	startRepl(cfg)
}
