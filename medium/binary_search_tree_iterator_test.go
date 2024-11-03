package main

import "testing"

// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
//
// BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
// boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
// int next() Moves the pointer to the right, then returns the number at the pointer.
// Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.
//
// You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.
//
// Example 1:
//
// Input
// ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// Output
// [null, 3, 7, true, 9, true, 15, true, 20, false]
//
// Explanation
// BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
// bSTIterator.next();    // return 3
// bSTIterator.next();    // return 7
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 9
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 15
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 20
// bSTIterator.hasNext(); // return False
func Test_binary_search_tree_iterator(t *testing.T) {
	root := TreeNode{
		Val: 7,
	}

	root.Left = &TreeNode{
		Val: 3,
	}

	root.Right = &TreeNode{
		Val: 15,
	}

	root.Right.Left = &TreeNode{
		Val: 9,
	}

	root.Right.Right = &TreeNode{
		Val: 20,
	}

	tests := []struct {
		root TreeNode
		want []interface{}
	}{
		{
			root: root,
			want: []interface{}{nil, 3, 7, 9, 15, 20},
		},
	}
	for _, tt := range tests {
		bstIterator := BSTIteratorConstructor(&tt.root)
		got := make([]interface{}, 0)
		got = append(got, nil)
		for bstIterator.HasNext() {
			got = append(got, bstIterator.Next())
		}
		if !equalTree(got, tt.want) {
			t.Errorf("Constructor(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}

func equalTree(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
