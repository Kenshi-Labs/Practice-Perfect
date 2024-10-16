package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test with an array where profit is possible
func TestMaxProfitPossible(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5 // Buy at 1, sell at 6
	result := maxProfit(prices)
	assert.Equal(t, expected, result, "Expected result should be equal to 5")
}

// Test with an array where no profit is possible (descending order)
func TestMaxProfitNoProfit(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0 // No profit can be made
	result := maxProfit(prices)
	assert.Equal(t, expected, result, "Expected result should be equal to 0")
}

// Test with an array containing duplicate values
func TestMaxProfitWithDuplicates(t *testing.T) {
	prices := []int{1, 1, 1, 1, 1}
	expected := 0 // No profit can be made
	result := maxProfit(prices)
	assert.Equal(t, expected, result, "Expected result should be equal to 0")
}

// Test with an array of length 2
func TestMaxProfitLengthTwo(t *testing.T) {
	prices := []int{1, 5}
	expected := 4 // Buy at 1, sell at 5
	result := maxProfit(prices)
	assert.Equal(t, expected, result, "Expected result should be equal to 4")
}