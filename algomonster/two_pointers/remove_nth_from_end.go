package two_pointers

// Remove N-th Node from End of Linked List
// Problem
// Given the head of a linked list and an integer n, remove the n-th node from the end of
// the list and return the head of the modified list.

// Input:

// head: the head node of a singly linked list
// n: an integer representing the position from the end (1-indexed)
// Output:

// Return the head of the modified linked list
// Constraints:

// The list has at least 1 node
// n is a valid position from the end (1 <= n <= length of list)
// Examples:

// Example 1:

// Input: head = [1, 2, 3, 4], n = 1
// Output: [1, 2, 3]
// Explanation: Remove the 1st node from the end (value 4)

// Example 2:

// Input: head = [1, 2, 3, 4], n = 2
// Output: [1, 2, 4]
// Explanation: Remove the 2nd node from the end (value 3)

// Example 3:

// Input: head = [1, 2, 3, 4], n = 4
// Output: [2, 3, 4]
// Explanation: Remove the 4th node from the end (the head, value 1)

func removeNthFromEnd(head *Node, n int) *Node {

	// use two pointers: left nad right
	// have the right move on each iteration,
	// while left waits until n iterations.
	// keep track of a previous pointer for left.
	// when we get to the end, have

	left, right := head, head
	var prev *Node
	counter := 1 - n

	for right.next != nil {
		if counter >= 0 {
			prev = left
			left = left.next
		}
		right = right.next
		counter++
	}

	// if prev was set, that means that the left
	// pointer started moving, and that we should
	// remove the node at the left pointer.
	if prev != nil {
		prev.next = left.next

	} else {
		// otherwise, we have to update head
		head = left.next
	}

	return head
}
