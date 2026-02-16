package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon to inspect")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon not found in pokedex: %s", pokemonName)
	}
	fmt.Printf("Name: %s\nheight: %d\nweight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	return nil
}
