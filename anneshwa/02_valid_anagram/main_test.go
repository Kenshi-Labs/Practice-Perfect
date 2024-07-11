package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.True(t, IsAnagram("anagram", "nagaram"))
	assert.True(t, IsAnagram("world hello", "hello world"))
	assert.False(t, IsAnagram("rat", "car"))
}