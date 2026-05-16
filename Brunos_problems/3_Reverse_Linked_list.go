package bruno

// 3. Reverse Linked List (Linked Lists)
// Problem: Reverse a singly linked list. Return the new head of the reversed list. (Assume a basic ListNode structure is used).

// Example:

// Input: 1 -> 2 -> 3 -> 4 -> 5 -> NULL

// Output: 5 -> 4 -> 3 -> 2 -> 1 -> NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(ll *ListNode) *ListNode {

	var prev *ListNode = nil
	current := ll
	var next *ListNode = current.Next

	if current.Next == nil {
		return ll
	}

	for {
		if next == nil {
			break
		}

		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// listFromSlice builds a singly linked list from a slice and returns the head.
// Returns nil for an empty slice.
func listFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// sliceFromList collects the values of a linked list into a slice.
// Returns an empty slice for a nil head.
func sliceFromList(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
