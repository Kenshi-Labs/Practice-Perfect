package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"", "", true},
		{"abc", "abcd", false},
	}

	for _, tt := range tests {
		t.Run(tt.s+"-"+tt.t, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
