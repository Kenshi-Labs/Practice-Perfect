package main

import (
	"strings"
)

// RemoveVowels removes all vowels (a, e, i, o, u) from the input string.
func RemoveVowels(s string) string {
	// Create a map to track vowels
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Build a new string without vowels
	var result strings.Builder
	for _, char := range s {
		if !vowels[char] {
			result.WriteRune(char)
		}
	}

	return result.String()
}
