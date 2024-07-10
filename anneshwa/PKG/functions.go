package functions

import (
	"strings"
	"unicode"
)

// PROBLEM 1: ReverseString reverses a string
func ReverseString(str string) string {
	// Convert string to a slice of runes
	runes := []rune(str)

	// Initialize two pointers for swapping characters
	left, right := 0, len(runes)-1

	// Iterate until the pointers meet in the middle
	for left < right {
		// Swap characters at left and right pointers
		runes[left], runes[right] = runes[right], runes[left]

		// Move the pointers towards the center
		left++
		right--
	}

	// Convert the rune slice back to a string and return
	return string(runes)
}

// PROBLEM 2: IsAnagram checks if two strings are anagrams
func IsAnagram(s string, t string) bool {
	// Remove spaces from strings
	s = removeSpaces(s)
	t = removeSpaces(t)

	// Early exit if lengths are different
	if len(s) != len(t) {
		return false
	}

	// Create maps to count character frequencies
	countS := make(map[rune]int)
	countT := make(map[rune]int)

	// Populate countS with frequencies of characters in s
	countCharacters(s, countS)

	// Populate countT with frequencies of characters in t
	countCharacters(t, countT)

	// Compare character frequencies
	for char, count := range countS {
		if countT[char] != count {
			return false
		}
	}

	return true
}

// Helper function to remove spaces from a string
func removeSpaces(s string) string {
	var result strings.Builder
	for _, char := range s {
		if !unicode.IsSpace(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// countCharacters counts character frequencies in s and updates countMap
func countCharacters(s string, countMap map[rune]int) {
	for _, char := range s {
		// Count every character, including spaces
		countMap[char]++
	}
}

// PROBLEM 3: TwoSum finds indices of two numbers that add up to a target
func TwoSum(nums []int, target int) []int {
	// Iterate over each number in the array
	for i := 0; i < len(nums); i++ {
		// Check for each number, if there is another number that sums up to the target
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				// Return the indices of the two numbers
				return []int{i, j}
			}
		}
	}

	// If no such pair is found, return an empty slice
	return nil
}

// PROBLEM 4: RemoveDuplicates function removes duplicates from a sorted array in-place
// and returns the new length.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}
	return uniqueIndex + 1
}

// LongestCommonPrefix function finds the longest common prefix string amongst an array of strings
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(str) < len(prefix) || prefix != str[:len(prefix)] {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
