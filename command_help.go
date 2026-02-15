package main

import (
	"fmt"
)

func helpCommand(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}
