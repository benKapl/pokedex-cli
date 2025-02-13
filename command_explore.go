package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandExplore(cfg *config, locationName string) error {
	if locationName == "" {
		return errors.New("you must specify a location")
	}

	fmt.Printf("Exploring %s...\n", locationName)

	exploredLocation, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if exploredLocation.Names == nil && err != nil {
		return errors.New("this location does not exist")
	}

	// Check for different errors than non existent location in Pokeapi
	if err != nil && !strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
		return errors.New(err.Error())
	}

	pokemons := exploredLocation.PokemonEncounters
	if len(pokemons) == 0 {
		fmt.Println("No pokemon found :(")
		return nil
	}

	fmt.Println("Found pokemon:")

	for _, p := range pokemons {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil
}
