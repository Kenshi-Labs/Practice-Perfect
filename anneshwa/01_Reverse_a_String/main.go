package main

import (
	"fmt"
 )

func main() {
	// Testing ReverseString
	fmt.Println("Testing ReverseString:")
	fmt.Println(ReverseString("hello"))           // Output: "olleh"
	fmt.Println(ReverseString(""))                // Output: ""
	fmt.Println(ReverseString("  hello world  ")) // Output: "  dlrow olleh  "
	fmt.Println(ReverseString("racecar"))         // Output: "racecar"

}

// PROBLEM 1: ReverseString reverses a string
func ReverseString(str string) string {
	// Convert string to a slice of runes
	runes := []rune(str)

	// Initialize two pointers for swapping characters
	left, right := 0, len(runes)-1

	// Iterate until the pointers meet in the middle
	for left < right {
		// Swap characters at left and right pointers
		runes[left], runes[right] = runes[right], runes[left]

		// Move the pointers towards the center
		left++
		right--
	}

	// Convert the rune slice back to a string and return
	return string(runes)
}