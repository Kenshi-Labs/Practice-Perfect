package main

import "fmt"

// LongestCommonPrefix finds the longest common prefix among the given strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// Check characters of the prefix against each string
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				// Update the prefix to the matched part
				prefix = prefix[:j]
				break
			}
		}
		// If no common prefix, return early
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(LongestCommonPrefix(strs)) // Output: "fl"
}
