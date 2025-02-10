package main

func main() {
	config := config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
	}

	startRepl(&config)
}
