package main

import "fmt"

func main() {
	data := "[{(()}]"
	fmt.Println(data, CheckParentheses(data))
}
