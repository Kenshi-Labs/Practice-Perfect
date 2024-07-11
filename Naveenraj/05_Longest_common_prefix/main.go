package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"flower", "flow", "flight"}

	fmt.Println(Longest_common_prefix(input))
}

func Longest_common_prefix(arr []string) string {
	output := ""

	smallWord := Find_a_smallword(arr)
	if len(smallWord) == 0 {
		return ""
	}

	checker := string(smallWord[0])
	for i := range smallWord {
		count := 0
		for _, val := range arr {
			if strings.HasPrefix(val, checker) {
				count++
			}
		}

		if count == len(arr) {
			output += string(smallWord[i])
		} else {
			break
		}

		if i < len(smallWord)-1 {
			checker += string(smallWord[i+1])
		}
	}
	return output
}

func Find_a_smallword(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	output := arr[0]

	for _, val := range arr {
		if len(val) < len(output) {
			output = val
		}
	}

	return output
}
