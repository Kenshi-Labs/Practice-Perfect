package main

import "fmt"

func main() {
	// Example usage: reverse
	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString(""))
	fmt.Println(reverseString("hello world"))
	fmt.Println(reverseString("racecar"))

	// Example usage: anagram
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false

	// Example usage: two_sum
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1, 2]
}
