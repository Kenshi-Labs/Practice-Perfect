package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "olleh", ReverseString("hello"))
	assert.Equal(t, "", ReverseString(""))
	assert.Equal(t, "  dlrow olleh  ", ReverseString("  hello world  "))
	assert.Equal(t, "racecar", ReverseString("racecar"))
}