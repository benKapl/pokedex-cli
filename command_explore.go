package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, locationName string) error {
	if locationName == "" {
		return errors.New("you must specify a location")
	}

	exploredLocation, _ := cfg.pokeapiClient.ExploreLocation(locationName)
	fmt.Printf("%+v\n", exploredLocation)
	return nil

	// TO DO :
	// - create the base function
	// - manage non existent locations

	// - cache management in internal pokeapi
	// - add test to repl_test.go (startRepl, mutliple parameters)

}
