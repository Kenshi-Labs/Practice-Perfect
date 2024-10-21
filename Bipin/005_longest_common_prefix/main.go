package main

import "fmt"

func main() {
	// Example usage and testing
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Input:", strs)
	fmt.Println("Longest common prefix:", longestCommonPrefix(strs))

	// Additional test cases
	testCases := [][]string{
		{"dog", "racecar", "car"},
		{"hello", "world", ""},
		{""},
		{"only"},
	}

	for _, tc := range testCases {
		fmt.Println("\nInput:", tc)
		fmt.Println("Longest common prefix:", longestCommonPrefix(tc))
	}
}
