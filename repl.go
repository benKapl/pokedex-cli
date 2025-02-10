package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := commandName.callback(cfg)
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
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
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
			description: "Display the names of 20 locations in the Pokemon World",
			callback:    commandMap,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splits := strings.Fields(text)
	return splits
}
