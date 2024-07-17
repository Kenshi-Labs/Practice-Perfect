package main

import "fmt"

func main() {
	input := ""
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

	// first method - failed in test case - {[]}

	// for i := 0; i < len(input); i = i + 2 {
	// 	expected := listOfBrackets[string(input[i])]
	// 	if string(input[i+1]) != expected {
	// 		return false
	// 	}
	// }

	// second method - failed in test case - ([)]

	// openMap := map[string]int{}
	// closeMap := map[string]int{}
	// for i := 0; i < len(input); i++ {
	// 	if IsOpen_Parentheses(string(input[i])) {
	// 		val, exist := openMap[string(input[i])]
	// 		if exist {
	// 			openMap[listOfBrackets[string(input[i])]] = val + 1
	// 		} else {
	// 			openMap[listOfBrackets[string(input[i])]] = 1
	// 		}
	// 	} else {
	// 		val, exist := closeMap[string(input[i])]
	// 		if exist {
	// 			closeMap[string(input[i])] = val + 1
	// 		} else {
	// 			closeMap[string(input[i])] = 1
	// 		}
	// 	}
	// }

	// for key, countInOpenMap := range openMap {
	// 	countInCloseMap := closeMap[key]
	// 	if countInCloseMap != countInOpenMap {
	// 		return false
	// 	}
	// }

	// return true

	//third method

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
