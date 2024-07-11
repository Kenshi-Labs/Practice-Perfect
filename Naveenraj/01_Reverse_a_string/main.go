package main

import "fmt"

func main(){
	input := "hello"
	reversedString := ReverseString(input)
	fmt.Println(reversedString)
}

func ReverseString(input string) string {
	output := ""
	for i := len(input)-1; i >=0; i--{
		output += string(input[i])
	}
	return output
}