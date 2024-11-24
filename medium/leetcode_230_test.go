package main

import "testing"

// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
//
//
//
//Example 1:
//
//
//Input: root = [3,1,4,null,2], k = 1
//Output: 1
//Example 2:
//
//
//Input: root = [5,3,6,2,4,null,null,1], k = 3
//Output: 3
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func TestKthSmallest(t *testing.T) {
	cases := []struct {
		root *TreeNode
		k    int
		want int
	}{
		//{
		//	root: &TreeNode{
		//		Val: 3,
		//		Left: &TreeNode{
		//			Val: 1,
		//			Right: &TreeNode{
		//				Val: 2,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 4,
		//		},
		//	},
		//	k:    1,
		//	want: 1,
		//},
		//{
		//	root: &TreeNode{
		//		Val: 5,
		//		Left: &TreeNode{
		//			Val: 3,
		//			Left: &TreeNode{
		//				Val: 2,
		//				Left: &TreeNode{
		//					Val: 1,
		//				},
		//			},
		//			Right: &TreeNode{
		//				Val: 4,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 6,
		//		},
		//	},
		//	k:    3,
		//	want: 3,
		//},
		//{
		//	root: &TreeNode{
		//		Val:  1,
		//		Left: nil,
		//		Right: &TreeNode{
		//			Val: 2,
		//		},
		//	},
		//	k:    2,
		//	want: 2,
		//},
		//{
		//	root: &TreeNode{
		//		Val: 4,
		//		Left: &TreeNode{
		//			Val: 2,
		//			Left: &TreeNode{
		//				Val: 3,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 5,
		//		},
		//	},
		//	k:    1,
		//	want: 2,
		//},
		// root: [2,1,3], k = 3
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			k:    3,
			want: 3,
		},
	}
	for _, c := range cases {
		if got := kthSmallest(c.root, c.k); got != c.want {
			t.Errorf("kthSmallest(%v, %v) = %v, want %v", c.root, c.k, got, c.want)
		}
	}
}
