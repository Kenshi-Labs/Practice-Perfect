// Unit tests
package main

import "fmt"

func main() {
	// Test 1: Two non-empty sorted lists
	l1 := createLinkedList([]int{1, 2, 4})
	l2 := createLinkedList([]int{1, 3, 4})
	mergedList := mergeTwoLists(l1, l2)
	fmt.Print("Merged List 1: ")
	printLinkedList(mergedList)

	// Test 2: One empty list, one non-empty list
	l1 = createLinkedList([]int{})
	l2 = createLinkedList([]int{0})
	mergedList = mergeTwoLists(l1, l2)
	fmt.Print("Merged List 2: ")
	printLinkedList(mergedList)

	// Test 3: Two empty lists
	l1 = createLinkedList([]int{})
	l2 = createLinkedList([]int{})
	mergedList = mergeTwoLists(l1, l2)
	fmt.Print("Merged List 3: ")
	printLinkedList(mergedList) // Should print nothing

	// Test 4: Lists of different lengths
	l1 = createLinkedList([]int{1, 2, 5, 6})
	l2 = createLinkedList([]int{1, 3, 4})
	mergedList = mergeTwoLists(l1, l2)
	fmt.Print("Merged List 4: ")
	printLinkedList(mergedList)
}
