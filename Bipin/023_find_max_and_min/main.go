package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 8, 2, 15}
	max, min := FindMaxMin(arr)
	fmt.Printf("Max = %d, Min = %d\n", max, min)
}
