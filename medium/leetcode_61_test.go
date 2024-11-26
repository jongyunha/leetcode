package main

import "testing"

// Given the head of a linked list, rotate the list to the right by k places.
//
// Example 1:
//
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
// Example 2:
//
// Input: head = [0,1,2], k = 4
// Output: [2,0,1]
//
// Constraints:
//
// The number of nodes in the list is in the range [0, 500].
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109
func TestRotateList(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		k        int
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:        2,
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			head:     &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			k:        4,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			k:        0,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		// head = [1,2,3,4,5], k = 3, expected = [3,4,5,1,2]
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:        3,
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}},
		},
	}

	for _, tc := range testCases {
		if res := rotateRight(tc.head, tc.k); !isListEqual(res, tc.expected) {
			t.Fatalf("expected %v but got %v", tc.expected, res)
		}
	}
}

func isListEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
