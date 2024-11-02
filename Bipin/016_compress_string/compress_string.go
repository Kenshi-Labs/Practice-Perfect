package main

import (
	"strconv"
	"strings"
)

// compressString performs string compression in a single pass using strings.Builder.
func CompressString(s string) string {
	// Edge case: If the string is empty, return it
	if len(s) == 0 {
		return s
	}

	// Using a strings.Builder for efficient string concatenation
	var builder strings.Builder
	countConsecutive := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			countConsecutive++
		} else {
			// Append character and count to the builder
			builder.WriteByte(s[i-1])
			builder.WriteString(strconv.Itoa(countConsecutive))
			countConsecutive = 1
		}
	}

	// Append the last character and its count
	builder.WriteByte(s[len(s)-1])
	builder.WriteString(strconv.Itoa(countConsecutive))

	// Convert the builder content to string
	compressed := builder.String()

	// If the compressed string is not smaller, return the original
	if len(compressed) >= len(s) {
		return s
	}

	return compressed
}
