package main

import (
	"container/list"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// row-order traversal of a tree, using Go's "container/ist" package
func rowOrderTreeTraversal(tree *TreeNode) []int {
	queue := list.New()

	rowOrder := []int{}

	// ut the head of the tree into the queue
	if tree != nil {
		queue.PushBack(tree)
	}
	for {

		// deque
		front := queue.Front() // pull the current fornt item from the queue
		if front != nil {
			queue.Remove(front)
		} else {
			break
		}

		currentNode := front.Value.(*TreeNode)

		// append the row-order slice
		rowOrder = append(rowOrder, currentNode.value)

		// and tree nodes if they exist
		if currentNode.left != nil {
			queue.PushBack(currentNode.left)
		}
		if currentNode.right != nil {
			queue.PushBack(currentNode.right)
		}

	}

	return rowOrder
}

func rowOrderTreeTraversalSlice(tree *TreeNode) []int {

	rowOrder := []int{}
	queue := []*TreeNode{}

	if tree == nil {
		return rowOrder
	}

	// put the first node into queue
	queue = append(queue, tree)

	for {

		// termination
		if len(queue) == 0 {
			break
		}

		// dequeue
		currentNode := queue[0]
		queue = queue[1:]
		rowOrder = append(rowOrder, currentNode.value)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}

	return rowOrder

}
