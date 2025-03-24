package main

import "testing"

// Given the root of a binary tree, return the most frequent subtree sum. If there is a tie, return all the values with the highest frequency in any order.
//
// The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself).
//
// Example 1:
//
// Input: root = [5,2,-3]
// Output: [2,-3,4]
// Example 2:
//
// Input: root = [5,2,-5]
// Output: [2]
func TestFindFrequentTreeSum(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}
	expected := []int{2, -3, 4}
	if result := findFrequentTreeSum(root); !equal(result, expected) {
		t.Errorf("findFrequentTreeSum(%v) = %v, expected %v", root, result, expected)
	}

	root = &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -5}}
	expected = []int{2}
	if result := findFrequentTreeSum(root); !equal(result, expected) {
		t.Errorf("findFrequentTreeSum(%v) = %v, expected %v", root, result, expected)
	}
}
