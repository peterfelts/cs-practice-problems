package main

import (
	"testing"
)

func Test_ReverseLinkedList(t *testing.T) {
	// Input: 1 -> 2 -> 3 -> 4 -> 5 -> NULL
	// Output: 5 -> 4 -> 3 -> 2 -> 1 -> NULL

	var node *Node = nil
	for i := 5; i > 0; i-- {
		node = &Node{i, node}
	}

	// Act
	reverseLinkedList(node)

}
