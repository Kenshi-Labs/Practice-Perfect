package main

import "testing"

// Test with a valid string of parentheses
func TestValidParentheses(t *testing.T) {
    input := "()[]{}"
    result := CheckParentheses(input)
    expected := true
    if result != expected {
        t.Errorf("TestValidParentheses failed: expected %v, got %v", expected, result)
    }
}

// Test with an invalid string of parentheses
func TestInvalidParentheses(t *testing.T) {
    input := "(]"
    result := CheckParentheses(input)
    expected := false
    if result != expected {
        t.Errorf("TestInvalidParentheses failed: expected %v, got %v", expected, result)
    }
}

// Test with an empty string
func TestEmptyString(t *testing.T) {
    input := ""
    result := CheckParentheses(input)
    expected := true
    if result != expected {
        t.Errorf("TestEmptyString failed: expected %v, got %v", expected, result)
    }
}

// Test with a string containing only opening brackets
func TestOnlyOpeningBrackets(t *testing.T) {
    input := "((("
    result := CheckParentheses(input)
    expected := false
    if result != expected {
        t.Errorf("TestOnlyOpeningBrackets failed: expected %v, got %v", expected, result)
    }
}

// Test with a complex valid string
func TestComplexValidString(t *testing.T) {
    input := "{[()()]}"
    result := CheckParentheses(input)
    expected := true
    if result != expected {
        t.Errorf("TestComplexValidString failed: expected %v, got %v", expected, result)
    }
}

// Test with a complex invalid string
func TestComplexInvalidString(t *testing.T) {
    input := "{[(])}"
    result := CheckParentheses(input)
    expected := false
    if result != expected {
        t.Errorf("TestComplexInvalidString failed: expected %v, got %v", expected, result)
    }
}

func TestMismatchedClosingParentheses(t *testing.T) {
    input := "())"
    result := CheckParentheses(input)
    expected := false
    if result != expected {
        t.Errorf("TestMismatchedClosingParentheses failed: expected %v, got %v", expected, result)
    }
}

// Test with excessive closing brackets
func TestExcessiveClosingBrackets(t *testing.T) {
    input := "(())}"
    result := CheckParentheses(input)
    expected := false
    if result != expected {
        t.Errorf("TestExcessiveClosingBrackets failed: expected %v, got %v", expected, result)
    }
}