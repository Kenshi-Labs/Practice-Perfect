package main

import "fmt"

// Stack struct represents a stack data structure,
// which holds a slice of integers to store the items.
type Stack struct {
	items []int
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop removes the top element from the stack.
// If the stack is empty, it does nothing.
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	// Removes the last element by slicing the stack
	s.items = s.items[:len(s.items)-1]
}

// Top returns the top element of the stack without removing it.
// If the stack is empty, it returns an error.
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	// Return the last element in the slice, which is the top of the stack
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// CheckParentheses checks if the input string has valid parentheses.
// It uses a stack to match every opening bracket with the corresponding closing bracket.
func CheckParentheses(input string) bool {
	var s Stack

	// Iterate through each character in the input string
	for _, ch := range input {

		// Push opening brackets onto the stack
		if ch == '[' || ch == '{' || ch == '(' {
			s.Push(int(ch))

			// For closing brackets, check if they match the top of the stack
		} else if ch == ']' || ch == '}' || ch == ')' {
			// Get the top element of the stack
			char, err := s.Top()
			if err != nil {
				// If stack is empty, it means no matching opening bracket
				return false
			}

			// Ensure the top of the stack has the corresponding opening bracket
			if (ch == ']' && char != '[') || (ch == '}' && char != '{') || (ch == ')' && char != '(') {
				return false
			}

			// Pop the matching opening bracket from the stack
			s.Pop()
		}
	}

	// If the stack is empty at the end, all brackets were properly matched
	return s.IsEmpty()
}
