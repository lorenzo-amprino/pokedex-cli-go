package main

import (
	"fmt"
)

func commandMapF(cfg *config, args []string) error {

	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	cfg.nextLocationsURL = &res.Next
	cfg.prevLocationsURL = &res.Previous

	return nil

}

func commandMapB(cfg *config, args []string) error {

	res, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	cfg.nextLocationsURL = &res.Next
	cfg.prevLocationsURL = &res.Previous

	return nil

}
