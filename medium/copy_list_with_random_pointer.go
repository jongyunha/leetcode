package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	dummy := &Node{}
	dummyPointer := dummy
	originalPointer := head

	for originalPointer != nil {
		newNode := &Node{Val: originalPointer.Val}
		nodeMap[originalPointer] = newNode
		dummyPointer.Next = newNode
		dummyPointer = dummyPointer.Next
		originalPointer = originalPointer.Next
	}

	originalPointer = head
	dummyPointer = dummy.Next

	for originalPointer != nil {
		if originalPointer.Random != nil {
			randomNode := nodeMap[originalPointer.Random]

			dummyPointer.Random = randomNode
		}

		originalPointer = originalPointer.Next
		dummyPointer = dummyPointer.Next
	}
	return dummy.Next
}
