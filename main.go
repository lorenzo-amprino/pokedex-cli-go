package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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
