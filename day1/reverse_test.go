package main

import (
	"testing"
)

// go test -v
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Non-empty string", "hello", "olleh"},
		{"Empty string", "", ""},
		{"String with spaces", "hello world", "dlrow olleh"},
		{"Palindrome", "racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, result)
			}
		})
	}
}
