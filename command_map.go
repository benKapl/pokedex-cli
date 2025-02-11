package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/benKapl/pokedex-cli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL != nil {
		return errors.New("you're on the first page")
	}

	locationsResp := pokeapi.RespShallowLocations{}

	// Get cache data if exists
	if data, exists := cfg.cache.Get(*cfg.prevLocationsURL); exists {
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return err
		}

	} else {
		locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}

	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
