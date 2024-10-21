package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello, World! u"
	fmt.Println(CountVowel(input))
}

func CountVowel(str string) int {
	vowel := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	count := 0
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		if _, pre := vowel[string(str[i])]; pre {
			count++
		}
	}

	return count
}
