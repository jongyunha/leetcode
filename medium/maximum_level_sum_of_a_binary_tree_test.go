package main

import "testing"

func TestMaximumLevelSumOfABinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: -8,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	got := maxLevelSum(root)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
