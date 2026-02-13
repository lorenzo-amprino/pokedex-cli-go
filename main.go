package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		fmt.Println("Your command was: " + inputCleaned[0])
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
