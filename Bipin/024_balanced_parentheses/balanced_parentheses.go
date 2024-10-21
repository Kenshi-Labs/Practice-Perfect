package main

// IsValid checks if the input string has balanced parentheses.
func IsValid(s string) bool {
	stack := []rune{} // Stack to hold opening brackets
	bracketMap := map[rune]rune{
		')': '(', // Map closing brackets to their corresponding opening brackets
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if _, exists := bracketMap[char]; exists { // If the character is a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				// Stack is empty or top of the stack is not the matching opening bracket
				return false
			}
			stack = stack[:len(stack)-1] // Pop the top of the stack
		} else {
			// If it's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}
	return len(stack) == 0 // If the stack is empty, all brackets are balanced
}
