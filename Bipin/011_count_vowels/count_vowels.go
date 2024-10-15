package main

import (
	"strings"
	"unicode"
)

// CountVowels: Function to count vowels in a string
func CountVowels(input string) int {
	vowels := "aeiou"
	count := 0

	// Iterate over each character in the string
	for _, char := range input {
		// Convert the character to lowercase and check if it's a vowel
		if strings.ContainsRune(vowels, unicode.ToLower(char)) {
			count++
		}
	}
	return count
}
