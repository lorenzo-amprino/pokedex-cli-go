package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon to catch")
	}

	pokemonDetails, err := cfg.pokeapiClient.GetPokemonDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetails.Name)
	chance := rand.Intn(pokemonDetails.BaseExperience)

	if float32(chance) > float32(pokemonDetails.BaseExperience)*0.65 {
		fmt.Printf("%s caught and added to Pokedex!\n", pokemonDetails.Name)
		cfg.pokedex[pokemonDetails.Name] = pokemonDetails
	} else {
		fmt.Printf("%s run away\n", pokemonDetails.Name)
	}

	return nil

}
