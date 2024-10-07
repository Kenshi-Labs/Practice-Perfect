package main

import (
	"strings"
	"unicode"
)

// CapitalizeFirstLetter: Function to capitalize the first letter of each word
func CapitalizeFirstLetter(input string) string {
	// Split the input into words
	words := strings.Fields(input)
	for i, word := range words {
		// Capitalize the first letter and combine with the rest of the word
		if len(word) > 0 {
			words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
		}
	}
	// Join the words back into a single string
	return strings.Join(words, " ")
}
