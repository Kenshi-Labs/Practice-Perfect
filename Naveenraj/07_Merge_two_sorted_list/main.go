package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to create a linked list from a slice of integers.
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Function to print a linked list.
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// Function to merge two sorted linked lists.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	// Create the linked lists from the slices.
	list1 := createLinkedList([]int{})
	list2 := createLinkedList([]int{})

	// Merge the linked lists.
	mergedList := mergeTwoLists(list1, list2)

	// Print the merged linked list.
	printLinkedList(mergedList)
}
