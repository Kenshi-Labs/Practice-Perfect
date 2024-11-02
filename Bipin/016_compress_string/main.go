package main

import "fmt"

// Unit Tests
func testCompressString() {
	// Test 1: String that can be compressed
	if CompressString("aabcccccaaa") != "a2b1c5a3" {
		fmt.Println("Test 1 Failed")
	} else {
		fmt.Println("Test 1 Passed")
	}

	// Test 2: String that cannot be compressed (should return original)
	if CompressString("abcdef") != "abcdef" {
		fmt.Println("Test 2 Failed")
	} else {
		fmt.Println("Test 2 Passed")
	}

	// Test 3: Empty string
	if CompressString("") != "" {
		fmt.Println("Test 3 Failed")
	} else {
		fmt.Println("Test 3 Passed")
	}

	// Test 4: String with no repeated characters (should return original)
	if CompressString("abcd") != "abcd" {
		fmt.Println("Test 4 Failed")
	} else {
		fmt.Println("Test 4 Passed")
	}
}

func main() {
	// Run the unit tests
	testCompressString()
}
