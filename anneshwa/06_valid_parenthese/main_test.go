package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
    // Test with a valid string of parentheses
    assert.True(t, IsValid("()[]{}"))

    // Test with an invalid string of parentheses
    assert.False(t, IsValid("(]"))

    // Test with an empty string
    assert.True(t, IsValid(""))

    // Test with a string containing only opening brackets
    assert.False(t, IsValid("((("))

    // Additional tests
    assert.True(t, IsValid("{[()]}"))
    assert.False(t, IsValid("{[(])}"))
    assert.False(t, IsValid("{[}"))
}