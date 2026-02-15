package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	commandsConfig := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	startRepl(commandsConfig)
}
