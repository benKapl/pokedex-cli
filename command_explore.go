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

	exploredLocation, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if exploredLocation.Names == nil && err != nil {
		return errors.New("this location does not exist")
	}

	// Check for different errors than non existent location in Pokeapi
	if err != nil && !strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
		return errors.New(err.Error())
	}

	pokemons := exploredLocation.PokemonEncounters

	// fmt.Println(pokemons)

	for _, p := range pokemons {
		fmt.Println(p.Pokemon.Name)
	}
	return nil

	// TO DO :
	// - create the base function
	// - manage non existent locations

	// - cache management in internal pokeapi
	// - add test to repl_test.go (startRepl, mutliple parameters)

}
