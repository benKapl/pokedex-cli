package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/benKapl/pokedex-cli/internal/pokeapi"
	"github.com/benKapl/pokedex-cli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	cache            pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var param string

		if len(words) > 1 {
			param = words[1]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg, param)
			if err != nil {
				fmt.Printf("Error : %v\n", err)
			}
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List and details available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List pokemons in the specified area",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splits := strings.Fields(text)
	return splits
}
