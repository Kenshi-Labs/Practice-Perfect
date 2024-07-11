package main

import (
	"fmt"
	"strings"
	"unicode"
	
 )

func main() {
	// Testing IsAnagram
	fmt.Println("Testing IsAnagram:")
	fmt.Println(IsAnagram("worldhello", "hello world")) // Output: true 
	fmt.Println(IsAnagram("rat", "car"))                // Output: false
	fmt.Println(IsAnagram("listen", "silent"))          // Output: true
	fmt.Println()
}

// PROBLEM 2: IsAnagram checks if two strings are anagrams
func IsAnagram(s string, t string) bool {
	// Remove spaces from strings
	s = removeSpaces(s)
	t = removeSpaces(t)

	// Early exit if lengths are different
	if len(s) != len(t) {
		return false
	}

	// Create maps to count character frequencies
	countS := make(map[rune]int)
	countT := make(map[rune]int)

	// Populate countS with frequencies of characters in s
	countCharacters(s, countS)

	// Populate countT with frequencies of characters in t
	countCharacters(t, countT)

	// Compare character frequencies
	for char, count := range countS {
		if countT[char] != count {
			return false
		}
	}

	return true
}

// Helper function to remove spaces from a string
func removeSpaces(s string) string {
	var result strings.Builder
	for _, char := range s {
		if !unicode.IsSpace(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// countCharacters counts character frequencies in s and updates countMap
func countCharacters(s string, countMap map[rune]int) {
	for _, char := range s {
		// Count every character, including spaces
		countMap[char]++
	}
}