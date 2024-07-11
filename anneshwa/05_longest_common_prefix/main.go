package main

import(
	"fmt"
)

func main(){
		// Testing LongestCommonPrefix
		fmt.Println("Testing LongestCommonPrefix:")
		fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
		fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
		fmt.Println(LongestCommonPrefix([]string{"", "b", "c"}))               // Output: ""
		fmt.Println(LongestCommonPrefix([]string{"alone"}))                    // Output: "alone"
		fmt.Println()
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