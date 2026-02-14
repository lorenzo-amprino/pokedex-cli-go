package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	commandsConfig := newCommandsConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputCleaned := cleanInput(scanner.Text())

		if len(inputCleaned) == 0 {
			continue
		}

		command := inputCleaned[0]

		if cmd, exists := getCommands()[command]; exists {
			if err := cmd.callback(commandsConfig); err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
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
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays a list of location areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous list of location areas",
			callback:    mapbCommand,
		},
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	// Split the input into words, trimming spaces and punctuation
	var words []string
	trimmed := strings.Trim(strings.ToLower(text), " \t\n\r")
	words = strings.Fields(trimmed)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandsConfig) error
}
