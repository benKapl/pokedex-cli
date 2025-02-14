package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must specify a pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		if strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
			return errors.New("this pokemon does not exist")
		}
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if !isCaught(pokemon.BaseExperience) {
		fmt.Printf("%s escaped !\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught !\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.pokedex[pokemon.Name] = pokemon // add Pokemon to pokedex
	return nil
}

func isCaught(baseXp int) bool {
	// Calculate probability of catching based on experience and a exponential curve.
	// The highest the xp the more difficult to catch a pokemon
	difficultyFactor := 0.005
	catchProb := math.Exp(-difficultyFactor * float64(baseXp))
	return rand.Float64() < catchProb

}
