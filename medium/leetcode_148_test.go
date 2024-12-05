package main

import "testing"

// Given the head of a linked list, return the list after sorting it in ascending order.
// Example 1:
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]
// Example 2:
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
// Example 3:
// Input: head = []
// Output: []
func TestSortList(t *testing.T) {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head = sortList(head)

	if !isSortedNode(head) {
		t.Errorf("List is not sorted")
	}

	head = &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	head = sortList(head)
	if !isSortedNode(head) {
		t.Errorf("List is not sorted")
	}

	head = nil
	head = sortList(head)
	if head != nil {
		t.Errorf("List is not sorted")
	}
}

func isSortedNode(head *ListNode) bool {
	current := head
	for current != nil && current.Next != nil {
		if current.Val > current.Next.Val {
			return false
		}
		current = current.Next
	}
	return true
}
