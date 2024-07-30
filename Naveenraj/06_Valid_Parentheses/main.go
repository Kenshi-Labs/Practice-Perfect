package main

import "fmt"

func main() {
	var input string
	fmt.Println("Enter the Parentheses to check valid or not")
	fmt.Scanln(&input)
	fmt.Println(Valid_Parentheses(input))
}

func Valid_Parentheses(input string) bool {
	listOfBrackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	if len(input) == 0 {
		return false
	}

	if len(input)%2 != 0 {
		return false
	}

	stack := []string{}
	for _, val := range input {
		//append all open Parentheses in stack variable
		if IsOpen_Parentheses(string(val)) {
			stack = append(stack, string(val))
		} else {
			lastbracket := len(stack) - 1
			//if previous stack and coming input (val) is not equal returning false
			if len(stack) == 0 || stack[lastbracket] != listOfBrackets[string(val)] {
				return false
			}
			//pop the last element
			stack = stack[:len(stack)-1]
		}
	}

	//if all open Parentheses have close Parentheses then exist true
	return len(stack) == 0
}

func IsOpen_Parentheses(str string) bool {
	return str == "(" || str == "{" || str == "["
}
