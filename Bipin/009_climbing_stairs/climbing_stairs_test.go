package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	// Test with n = 1
	result := ClimbStairs(1)
	expected := 1
	if result != expected {
		t.Errorf("Test failed for n=1, expected %d, got %d", expected, result)
	}

	// Test with n = 2
	result = ClimbStairs(2)
	expected = 2
	if result != expected {
		t.Errorf("Test failed for n=2, expected %d, got %d", expected, result)
	}

	// Test with n = 3
	result = ClimbStairs(3)
	expected = 3
	if result != expected {
		t.Errorf("Test failed for n=3, expected %d, got %d", expected, result)
	}

	// Test with n = 4
	result = ClimbStairs(4)
	expected = 5
	if result != expected {
		t.Errorf("Test failed for n=4, expected %d, got %d", expected, result)
	}

	// Test with n = 5
	result = ClimbStairs(5)
	expected = 8
	if result != expected {
		t.Errorf("Test failed for n=5, expected %d, got %d", expected, result)
	}

	// Test with n = 0
	result = ClimbStairs(0)
	expected = 0
	if result != expected {
		t.Errorf("Test failed for n=0, expected %d, got %d", expected, result)
	}
}
