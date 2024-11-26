package main

import "testing"

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
//
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
//
// Example 1:
//
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.
// Example 2:
//
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
// Example 3:
//
// Input: root = [1,2], p = 1, q = 2
// Output: 1
func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	//p := root.Left
	//q := root.Right
	//
	//got := lowestCommonAncestor(root, p, q)
	//if got.Val != 3 {
	//	t.Errorf("lowestCommonAncestor() = %v; want 3", got.Val)
	//}

	// example 2
	p := root.Left
	q := root.Left.Right.Right

	got := lowestCommonAncestor(root, p, q)
	if got.Val != 5 {
		t.Errorf("lowestCommonAncestor() = %v; want 5", got.Val)
	}

}
