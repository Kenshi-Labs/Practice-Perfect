package main

import (
	"sort"
)

// isAnagram checks if string t is an anagram of string s.
func isAnagram(s string, t string) bool {
	// Early return if lengths are different, they cannot be anagrams.
	if len(s) != len(t) {
		return false
	}

	// Convert strings to slices of runes for comparison.
	sRunes := []rune(s)
	tRunes := []rune(t)

	// Sort both slices of runes.
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// Compare sorted slices rune by rune.
	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}
