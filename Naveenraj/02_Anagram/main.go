package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anag  ram"
	t := "nagaram"
	fmt.Println(IsAnagram(s, t))
}

func IsAnagram(s, t string) bool {
	s = strings.ReplaceAll(s, " ", "") //removing all space
	t = strings.ReplaceAll(t, " ", "") //removing all space

	arr1 := strings.Split(s, "") // converting string into array
	arr2 := strings.Split(t, "") // converting string into array

	sort.Strings(arr1) // sorting the array
	sort.Strings(arr2) // sorting the array

	str1 := strings.Join(arr1, "") // joining the sorted array
	str2 := strings.Join(arr2, "") // joining the sorted array

	return str1 == str2
}
