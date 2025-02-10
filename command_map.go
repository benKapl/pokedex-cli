package main

import (
	"fmt"

	"github.com/benKapl/pokeapi"
)

func commandMap(config *config) error {
	result, err := pokeapi.FetchLocations("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("config : %+v", *config)
	fmt.Printf("result : %+v", result)
	return nil

}
