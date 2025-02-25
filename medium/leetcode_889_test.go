package main

import "testing"

// Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary tree of distinct values and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.
//
// If there exist multiple answers, you can return any of them.
//
// Example 1:
//
// Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// Output: [1,2,3,4,5,6,7]
// Example 2:
//
// Input: preorder = [1], postorder = [1]
// Output: [1]
func TestConstructFromPrePost(t *testing.T) {
	testCases := []struct {
		preorder  []int
		postorder []int
		want      *TreeNode
	}{
		//{
		//	preorder:  []int{1, 2, 4, 5, 3, 6, 7},
		//	postorder: []int{4, 5, 2, 6, 7, 3, 1},
		//	want: &TreeNode{
		//		Val: 1,
		//		Left: &TreeNode{
		//			Val: 2,
		//			Left: &TreeNode{
		//				Val: 4,
		//			},
		//			Right: &TreeNode{
		//				Val: 5,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 3,
		//			Left: &TreeNode{
		//				Val: 6,
		//			},
		//			Right: &TreeNode{
		//				Val: 7,
		//			},
		//		},
		//	},
		//},
		//{
		//	preorder:  []int{1},
		//	postorder: []int{1},
		//	want: &TreeNode{
		//		Val: 1,
		//	},
		//},
		{
			preorder:  []int{2, 1, 3},
			postorder: []int{3, 1, 2},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	for _, tc := range testCases {
		got := constructFromPrePost(tc.preorder, tc.postorder)
		if !isSameTree(got, tc.want) {
			t.Errorf("preorder: %v, postorder: %v, got: %v, want: %v", tc.preorder, tc.postorder, got, tc.want)
		}
	}
}
