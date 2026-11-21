package two_pointers

// Middle of a Linked List
// Find the middle node of a linked list.

// Input: 0 1 2 3 4

// Output: 2

// If the number of nodes is even, then return the second middle node.

// Input: 0 1 2 3 4 5

// Output: 3

type Node struct {
	val  int
	next *Node
}

func middleOfLinkedList(head *Node) int {
	// WRITE YOUR BRILLIANT CODE HERE

	// Two pointers: slow and fast
	// slow increments by one
	// fast increments by two
	// when fast gets to the end, slow should be pointing
	// to the middle
	slow, fast := head, head

	for fast.next != nil {

		slow = slow.next

		// special case where we can't move the fast pointer
		// any more
		if fast.next.next == nil {
			return slow.val
		}
		fast = fast.next.next

	}

	return slow.val
}
