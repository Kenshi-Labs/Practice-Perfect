package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted linked lists and return it as a sorted list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Base cases: if either list is empty, return the other list
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var merged *ListNode

	// Recursively merge the two lists
	if l1.Val < l2.Val {
		merged = l1
		merged.Next = mergeTwoLists(l1.Next, l2)
	} else {
		merged = l2
		merged.Next = mergeTwoLists(l1, l2.Next)
	}

	return merged
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list back to a slice for easy comparison
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
