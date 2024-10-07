package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(0))
}

func climbStairs(n int) int {
	output := 0
	a, b := 0, 1
	for i := 0; i < n; i++ {
		output = a + b
		a = b
		b = output
	}

	return output
}
