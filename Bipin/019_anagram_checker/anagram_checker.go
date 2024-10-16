package main

import (
	"strings"
)

// AnagramChecker checks if two strings are anagrams using a single frequency map
func AnagramChecker(s1, s2 string) bool {
	// Remove spaces and convert both strings to lowercase
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	// If lengths are not equal, they can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// Use a single frequency map
	freqMap := make(map[rune]int)

	// Increase the count for s1 characters, decrease for s2 characters
	for _, char := range s1 {
		freqMap[char]++
	}
	for _, char := range s2 {
		freqMap[char]--
	}

	// Check if all values in the frequency map are zero
	for _, count := range freqMap {
		if count != 0 {
			return false
		}
	}

	return true
}
