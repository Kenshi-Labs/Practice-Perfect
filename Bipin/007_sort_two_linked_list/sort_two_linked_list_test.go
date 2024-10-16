package main

import (
	"reflect"
	"testing"
)

// Test case 1: Test with two non-empty sorted lists
func TestMergeTwoNonEmptySortedLists(t *testing.T) {
	l1 := createLinkedList([]int{1, 2, 4})
	l2 := createLinkedList([]int{1, 3, 4})
	expected := []int{1, 1, 2, 3, 4, 4}

	mergedList := mergeTwoLists(l1, l2)
	result := linkedListToSlice(mergedList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test case 2: Test with one empty list and one non-empty list
func TestMergeWithOneEmptyList(t *testing.T) {
	l1 := createLinkedList([]int{})
	l2 := createLinkedList([]int{0})
	expected := []int{0}

	mergedList := mergeTwoLists(l1, l2)
	result := linkedListToSlice(mergedList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test case 3: Test with two empty lists
func TestMergeWithTwoEmptyLists(t *testing.T) {
	l1 := createLinkedList([]int{})
	l2 := createLinkedList([]int{0})
	expected := []int{0}

	mergedList := mergeTwoLists(l1, l2)
	result := linkedListToSlice(mergedList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test case 4: Test with lists of different lengths
func TestMergeListsOfDifferentLengths(t *testing.T) {
	l1 := createLinkedList([]int{1, 2, 5, 6})
	l2 := createLinkedList([]int{1, 3, 4})
	expected := []int{1, 1, 2, 3, 4, 5, 6}

	mergedList := mergeTwoLists(l1, l2)
	result := linkedListToSlice(mergedList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
