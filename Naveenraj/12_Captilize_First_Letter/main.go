package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "First leTTeR of EACH Word"
	fmt.Println(CaptilizeFirstLetter(input))
}

func CaptilizeFirstLetter(str string) string {
	splited := strings.Split(str, " ")
	output := []string{}
	for _, val := range splited {
		if len(val) > 2 {
			output = append(output, strings.ToUpper(string(val[0]))+strings.ToLower(val[1:]))
		} else {
			output = append(output, strings.ToLower(val))
		}
	}

	return strings.Join(output, " ")
}
