package main

import "fmt"

func main() {
	// Running some sample test cases
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
	fmt.Println(LongestCommonPrefix([]string{"", "flow", "flight"}))       // Output: ""
	fmt.Println(LongestCommonPrefix([]string{"single"}))                   // Output: "single"
}
