package main

// strings.ToLower() to convert a string to lowercase 
// and strings.TrimSpace() to remove leading and trailing whitespace
// strings.Fields() function, which divides the string into substrings
// by removing any spaces, including newlines, and returns a slice of the resulting substrings

import (
	"strings"
)

func cleanInput(text string) []string {
	result := []string{}
	words := strings.Fields(text)
	for _, w := range words {
		lowered := strings.ToLower(w)
		result = append(result, lowered)
	}
	return result
}