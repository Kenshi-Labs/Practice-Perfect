package main

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
func longestCommonPrefix(strs []string) string {
	// If input array is empty, return an empty string
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	// Iterate through each string in the array
	for _, str := range strs[1:] {
		// Adjust prefix to match the current string
		for !startsWith(str, prefix) {
			// Remove the last character of prefix
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				// If prefix becomes empty, return immediately
				return ""
			}
		}
	}

	// Return the longest common prefix found
	return prefix
}

// startsWith checks if the string s starts with prefix.
func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
