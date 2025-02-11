package main

import (
	"time"

	"github.com/benKapl/pokedex-cli/internal/pokeapi"
	"github.com/benKapl/pokedex-cli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokeCache,
	}

	startRepl(cfg)
}
