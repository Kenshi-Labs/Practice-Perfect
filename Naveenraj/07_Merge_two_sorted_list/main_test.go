package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoSortList(t *testing.T) {
	testcases := []struct {
		arr1     []int
		arr2     []int
		expected []int
	}{
		{arr1: []int{1, 2, 4}, arr2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{arr1: []int{}, arr2: []int{1}, expected: []int{1}},
		{arr1: []int{}, arr2: []int{}, expected: []int{}},
		{arr1: []int{1, 2, 4}, arr2: []int{1, 3}, expected: []int{1, 1, 2, 3, 4}},
	}

	for i, val := range testcases {
		list1 := createLinkedList(val.arr1)
		list2 := createLinkedList(val.arr2)
		mergedList := mergeTwoLists(list1, list2)
		output := linkedListToSlice(mergedList)
		if !reflect.DeepEqual(output, val.expected) {
			t.Errorf("Error in the test case no %d ", i+1)
		} else {
			fmt.Printf("Test case %d passes \n", i+1)
		}
	}
}

func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
