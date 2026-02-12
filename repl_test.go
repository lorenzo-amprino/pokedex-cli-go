package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello, World!  ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "Go is great!",
			expected: []string{"go", "is", "great!"},
		},
		{
			input:    "  Multiple   spaces  ",
			expected: []string{"multiple", "spaces"},
		},
	}
	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("Expected %d words, got %d", len(c.expected), len(result))
			continue
		}
		for i, word := range result {
			if word != c.expected[i] {
				t.Errorf("Expected word '%s', got '%s'", c.expected[i], word)
			}
		}
	}
}
