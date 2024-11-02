package main

import "strings"

// reverseByWord takes a string 's' and returns a new string
// where the words in 's' are reversed in order.
func reverseByWord(s string) string {
	// Remove leading and trailing whitespace from the input string
	trimmedStr := strings.TrimSpace(s)

	// Split the trimmed string into a slice of words using space as the delimiter
	sp := strings.Split(trimmedStr, " ")

	// Swap elements from the start and end of the slice to reverse the order of words
	for i, j := 0, len(sp)-1; i < j; i, j = i+1, j-1 {
		sp[i], sp[j] = sp[j], sp[i] // Swap the words at indices i and j
	}

	// Join the reversed slice of words into a single string, separated by spaces
	return strings.Join(sp, " ")
}
