package main

import "fmt"

// Unit tests
func main() {
	// Test with 1 stair
	fmt.Println(ClimbStairs(1))

	// Test with 2 stair
	fmt.Println(ClimbStairs(2))

	// Test with 5 stair
	fmt.Println(ClimbStairs(5))

	// Test with 0 stair
	fmt.Println(ClimbStairs(0))
}
