package main

import (
	"fmt"

	functions "github.com/Kenshi-Labs/Practice-Perfect/pkg"
)

func main() {
	// Testing ReverseString
	fmt.Println("Testing ReverseString:")
	fmt.Println(functions.ReverseString("hello"))           // Output: "olleh"
	fmt.Println(functions.ReverseString(""))                // Output: ""
	fmt.Println(functions.ReverseString("  hello world  ")) // Output: "  dlrow olleh  "
	fmt.Println(functions.ReverseString("racecar"))         // Output: "racecar"
	fmt.Println()

	// Testing IsAnagram
	fmt.Println("Testing IsAnagram:")
	fmt.Println(functions.IsAnagram("worldhello", "hello world")) // Output: false (spaces matter)
	fmt.Println(functions.IsAnagram("rat", "car"))                // Output: false
	fmt.Println(functions.IsAnagram("listen", "silent"))          // Output: true
	fmt.Println()

	// Testing TwoSum
	fmt.Println("Testing TwoSum:")
	fmt.Println(functions.TwoSum([]int{2, 7, 11, 15}, 9)) // Output: [0 1]
	fmt.Println(functions.TwoSum([]int{3, 2, 4}, 6))      // Output: [1 2]
	fmt.Println()

	// Testing LongestCommonPrefix
	fmt.Println("Testing LongestCommonPrefix:")
	fmt.Println(functions.LongestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(functions.LongestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
	fmt.Println(functions.LongestCommonPrefix([]string{"", "b", "c"}))               // Output: ""
	fmt.Println(functions.LongestCommonPrefix([]string{"alone"}))                    // Output: "alone"
	fmt.Println()

	// Testing RemoveDuplicates
	fmt.Println("Testing RemoveDuplicates:")
	arr1 := []int{1, 1, 2}
	newLength1 := functions.RemoveDuplicates(arr1)
	fmt.Println(newLength1, arr1[:newLength1]) // Output: 2 [1 2]

	arr2 := []int{1, 2, 3}
	newLength2 := functions.RemoveDuplicates(arr2)
	fmt.Println(newLength2, arr2[:newLength2]) // Output: 3 [1 2 3]

	arr3 := []int{}
	newLength3 := functions.RemoveDuplicates(arr3)
	fmt.Println(newLength3, arr3[:newLength3]) // Output: 0 []

	arr4 := []int{2, 2, 2, 2}
	newLength4 := functions.RemoveDuplicates(arr4)
	fmt.Println(newLength4, arr4[:newLength4]) // Output: 1 [2]
}
