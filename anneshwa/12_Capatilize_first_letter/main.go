package main

import (
	"fmt"
	"strings"
)

// CapitalizeFirstLetter capitalizes the first letter of each word in a string.
func CapitalizeFirstLetter(input string) string {
	// Split the input string into words
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			// Capitalize the first letter of each word
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	// Join the words back into a single string
	return strings.Join(words, " ")
}

func main() {
	input := "examPLE input strIng"
	output := CapitalizeFirstLetter(input)
	fmt.Printf("Custom Input: \"%s\" -> Output: \"%s\"\n", input, output)
}
