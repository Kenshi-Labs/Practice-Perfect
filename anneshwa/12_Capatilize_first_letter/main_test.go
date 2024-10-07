package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	// Test with a sentence containing multiple words
	input := "the quick brown fox"
	expected := "The Quick Brown Fox"
	output := CapitalizeFirstLetter(input)
	assert.Equal(t, expected, output, "The output should capitalize the first letter of each word")

	// Test with a single word
	input = "example"
	expected = "Example"
	output = CapitalizeFirstLetter(input)
	assert.Equal(t, expected, output, "The output should capitalize the first letter of the single word")

	// Test with a string containing numbers and special characters
	input = "123 testing! @special characters"
	expected = "123 Testing! @special Characters"
	output = CapitalizeFirstLetter(input)
	assert.Equal(t, expected, output, "The output should correctly capitalize words while ignoring numbers and special characters")

	// Test with an empty string
	input = ""
	expected = ""
	output = CapitalizeFirstLetter(input)
	assert.Equal(t, expected, output, "The output of an empty string should also be an empty string")
}
