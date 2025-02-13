package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must specify a location")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.ExploreLocation(name)
	if location.Names == nil && err != nil {
		return errors.New("this location does not exist")
	}

	// Check for different errors than non existent location in Pokeapi
	if err != nil && !strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
		return errors.New(err.Error())
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}
	return nil
}
