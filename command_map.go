package main

import (
	"fmt"

	"github.com/benKapl/pokeapi"
)

func commandMap(config *config) error {
	res, err := pokeapi.FetchLocations(config.Next)
	if err != nil {
		return err
	}
	config.Next = res.Next // update Next
	config.Previous = res.Previous

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapB(config *config) error {
	if config.Previous != nil {

		res, err := pokeapi.FetchLocations(*config.Previous)
		if err != nil {
			return err
		}
		config.Next = res.Next // update Next
		config.Previous = res.Previous

		for _, result := range res.Results {
			fmt.Println(result.Name)
		}

		return nil
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
}
