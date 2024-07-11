package main

import "fmt"

func main() {
	// Example usage: anagram
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
}
