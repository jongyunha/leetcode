package main

import "testing"

func TestSumRootToLeafToNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	got := sumNumbers(root)
	want := 25
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	got = sumNumbers(root)
	want = 1026
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
