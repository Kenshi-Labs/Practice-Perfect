package main

import (
	"fmt"
)
func main(){
	fmt.Println("Testing Valid Parenthese:")
	fmt.Println(IsValid("(){}[]"))           // Output: true
}

func IsValid(s string) bool {
    // Map to hold the mappings of closing to opening brackets
    bracketMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    // Stack to keep track of opening brackets
    stack := []rune{}

    // Iterate over each character in the string
    for _, char := range s {
        // If it's a closing bracket
        if open, exists := bracketMap[char]; exists {
            // Check if the stack is not empty and the top of the stack matches the corresponding opening bracket
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            }
            // Pop the top of the stack
            stack = stack[:len(stack)-1]
        } else {
            // It's an opening bracket, push it onto the stack
            stack = append(stack, char)
        }
    }

    // If the stack is empty, all opening brackets have been matched
    return len(stack) == 0
}