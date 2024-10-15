package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("Common prefix", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		expected := "fl"
		result := LongestCommonPrefix(strs)
		assert.Equal(t, expected, result)
	})

	t.Run("No common prefix", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		expected := ""
		result := LongestCommonPrefix(strs)
		assert.Equal(t, expected, result)
	})

	t.Run("Array containing an empty string", func(t *testing.T) {
		strs := []string{"", "b", "c"}
		expected := ""
		result := LongestCommonPrefix(strs)
		assert.Equal(t, expected, result)
	})

	t.Run("Array containing only one string", func(t *testing.T) {
		strs := []string{"single"}
		expected := "single"
		result := LongestCommonPrefix(strs)
		assert.Equal(t, expected, result)
	})
}
