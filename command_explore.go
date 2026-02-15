package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location to explore")
	}
	location := args[0]
	fmt.Printf("Exploring %s\n", location)

	res, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
