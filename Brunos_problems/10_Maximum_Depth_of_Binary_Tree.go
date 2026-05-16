package bruno

// 10. Maximum Depth of Binary Tree (Binary Trees & Recursion)
// Problem: Given the root of a binary tree, return its maximum depth. The
// maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example (Conceptual):

// Input (Tree structure): A tree with root 3, left child 9, and right child
// 20 (which has children 15 and 7).

// Output: 3 (The path is 3 -> 20 -> 15 or 3 -> 20 -> 7)

type Node struct {
	value int
	left  *Node
	right *Node
}

func calculateMaximumDepthOfBinaryTree(root *Node) int {

	if root == nil {
		return 0
	}

	return calculateMaximumDepthOfBinaryTree_Impl(root, 1)

}

func calculateMaximumDepthOfBinaryTree_Impl(node *Node, depth int) int {

	depthLeft := depth
	depthRight := depth

	if node.left != nil {
		depthLeft = calculateMaximumDepthOfBinaryTree_Impl(node.left, depthLeft+1)
	}
	if node.right != nil {
		depthRight = calculateMaximumDepthOfBinaryTree_Impl(node.right, depthRight+1)
	}

	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight

}
