package main

// ClimbStairs calculates how many distinct ways you can climb to the top
func ClimbStairs(n int) int {
	// Handle base cases
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	// Initialize the first two steps
	first, second := 1, 2

	// Iterate from step 3 to n
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}
