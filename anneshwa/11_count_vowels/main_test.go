package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test with a string containing vowels
func TestCountVowelsWithVowels(t *testing.T) {
	input := "Hello, World!"
	expected := 3 // 'e', 'o', 'o'
	result := countVowels(input)
	assert.Equal(t, expected, result, "The number of vowels should be 3")
}

// Test with a string containing no vowels
func TestCountVowelsNoVowels(t *testing.T) {
	input := "rhythm"
	expected := 0 // No vowels
	result := countVowels(input)
	assert.Equal(t, expected, result, "The number of vowels should be 0")
}

// Test with an empty string
func TestCountVowelsEmptyString(t *testing.T) {
	input := ""
	expected := 0 // No characters
	result := countVowels(input)
	assert.Equal(t, expected, result, "The number of vowels should be 0")
}

// Test with a string containing only vowels
func TestCountVowelsOnlyVowels(t *testing.T) {
	input := "aeiouAEIOU"
	expected := 10 // 5 lowercase vowels + 5 uppercase vowels
	result := countVowels(input)
	assert.Equal(t, expected, result, "The number of vowels should be 10")
}