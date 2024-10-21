package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	t.Run("Compressible string", func(t *testing.T) {
		result := Compress("aabcccccaaa")
		assert.Equal(t, "a2b1c5a3", result)
	})

	t.Run("Non-compressible string", func(t *testing.T) {
		result := Compress("abc")
		assert.Equal(t, "abc", result)
	})

	t.Run("Empty string", func(t *testing.T) {
		result := Compress("")
		assert.Equal(t, "", result)
	})

	t.Run("String with no repeats", func(t *testing.T) {
		result := Compress("abcd")
		assert.Equal(t, "abcd", result)
	})
}
