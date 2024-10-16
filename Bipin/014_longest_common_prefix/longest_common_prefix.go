package main

// LongestCommonPrefix function finds the longest common prefix among an array of strings
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Take the first string as reference
	prefix := strs[0]

	// Compare the prefix with each string in the array
	for i := 1; i < len(strs); i++ {
		// Check the common prefix with the current string
		for len(prefix) > 0 && !startsWith(strs[i], prefix) {
			// Reduce the prefix by removing one character at a time
			prefix = prefix[:len(prefix)-1]
		}
		// If there's no common prefix, return an empty string
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// startsWith checks if the string s starts with prefix
func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}
