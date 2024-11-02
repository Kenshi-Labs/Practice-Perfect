package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairsWith1(t *testing.T) {
	n := 1
	expected := 1
	result := climbStairs(n)
	assert.Equal(t, expected, result, "Expected result should be equal to 1")
}

func TestClimbStairsWith2(t *testing.T) {
	n := 2
	expected := 2
	result := climbStairs(n)
	assert.Equal(t, expected, result, "Expected result should be equal to 2")
}

func TestClimbStairsWith5(t *testing.T) {
	n := 5
	expected := 8
	result := climbStairs(n)
	assert.Equal(t, expected, result, "Expected result should be equal to 8")
}

func TestClimbStairsWith0(t *testing.T) {
	n := 0
	expected := 1
	result := climbStairs(n)
	assert.Equal(t, expected, result, "Expected result should be equal to 1")
}
