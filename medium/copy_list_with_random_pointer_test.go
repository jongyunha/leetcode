package main

import "testing"

// Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
func Test_copy_list_with_random_pointer(t *testing.T) {
	// Create the original list with random pointers
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	// Expected output list
	expectedNode1 := &Node{Val: 7}
	expectedNode2 := &Node{Val: 13}
	expectedNode3 := &Node{Val: 11}
	expectedNode4 := &Node{Val: 10}
	expectedNode5 := &Node{Val: 1}

	expectedNode1.Next = expectedNode2
	expectedNode2.Next = expectedNode3
	expectedNode3.Next = expectedNode4
	expectedNode4.Next = expectedNode5

	expectedNode1.Random = nil
	expectedNode2.Random = expectedNode1
	expectedNode3.Random = expectedNode5
	expectedNode4.Random = expectedNode3
	expectedNode5.Random = expectedNode1

	// Call the function to copy the list
	copiedList := copyRandomList(node1)

	// Function to compare two lists
	compareLists := func(l1, l2 *Node) bool {
		for l1 != nil && l2 != nil {
			if l1.Val != l2.Val || (l1.Random != nil && l2.Random != nil && l1.Random.Val != l2.Random.Val) {
				return false
			}
			l1 = l1.Next
			l2 = l2.Next
		}
		return l1 == nil && l2 == nil
	}

	// Check if the copied list matches the expected list
	if !compareLists(copiedList, expectedNode1) {
		t.Errorf("copyRandomList() did not produce the expected output")
	}
}
