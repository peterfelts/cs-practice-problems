package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// Input: 1 -> 2 -> 3 -> 4 -> 5 -> NULL
// Output: 5 -> 4 -> 3 -> 2 -> 1 -> NULL
func reverseLinkedList(node *Node) *Node {

	printLinkedList(node)

	var previous *Node = nil
	var current *Node = node // starts with 1
	var next *Node = nil

	for {
		if current == nil {
			break
		}
		next = current.next     // 2, 3
		current.next = previous // nil <- 1 <- 2
		previous = current      // 1, 2
		current = next          // 2, 3

	}

	// print reverse linked list
	printLinkedList(previous)

	return previous
}

func printLinkedList(node *Node) {
	for {
		fmt.Printf("%d->", node.value)
		node = node.next
		if node == nil {
			fmt.Println()
			return
		}
	}
}
