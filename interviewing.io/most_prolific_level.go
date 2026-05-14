// To execute Go code, please declare a func main() in a package "main"

package interviewingio

import "fmt"

// # Most Prolific Level

// Given the root of a binary tree, return the most prolific level. The _prolificness_ of a level is the average number of children over all the nodes in that level.

// Return `-1` if the tree is empty. In case of a tie, return any of the most prolific levels.

// Example:

// Input:
//       O
//      /
//     O
//    / \
//   O   O
//  / \   \
// O   O   O

// Output: 1
// - Level 0 has prolificness 1
// - Level 1 has prolificness 2
// - Level 2 has prolificness 1.5
// - Level 3 has prolificness 0
// The most prolific level is 1, with a prolificness of 2.

// https://iio-beyond-ctci-images.s3.us-east-1.amazonaws.com/trees_fig18.png

// Constraints:

// - The number of nodes is at most `10^5`
// - The value at each node doesn't matter.

type Node struct {
	Value     int
	Left      *Node
	Right     *Node
	RowNumber int
}

func calculateProlificness(node *Node) int {

	if node == nil {
		return -1
	}

	queue := []*Node{node}
	rowCounts := []int{}

	prevRowNumber := -1
	rowSum := 0

	// Perform a BFS search of the tree. When a new row is detected,
	// record the length of the previous row
	for {

		rowSum++
		if len(queue) == 0 {
			// When we're at the end of the final row, we also need to
			// record the row cown
			rowCounts = append(rowCounts, rowSum)
			break
		}

		// Pop the queue
		node := queue[0]
		queue = queue[1:]

		// detect when we've reached another row
		if node.RowNumber > prevRowNumber {
			if prevRowNumber > -1 {
				rowCounts = append(rowCounts, rowSum)
			}
			prevRowNumber = node.RowNumber
			rowSum = 0
		}

		// Add left and right nodes to the queue
		if node.Left != nil {
			node.Left.RowNumber = node.RowNumber + 1
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			node.Right.RowNumber = node.RowNumber + 1
			queue = append(queue, node.Right)
		}
	}

	// iterate through the row counts, calculate prolificness and identify the
	// row with the maximum prolificness
	maxProlificness := float64(0)
	maxProlificnessRow := 0
	for i := 0; i < len(rowCounts)-1; i++ {
		prolificness := float64(rowCounts[i+1]) / float64(rowCounts[i])
		if prolificness > float64(maxProlificness) {
			maxProlificness = prolificness
			maxProlificnessRow = i
		}
	}

	return maxProlificnessRow
}

func BuildTree(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*Node, len(values))

	for i, v := range values {
		nodes[i] = &Node{Value: v}
	}

	for i := 0; i < len(values); i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < len(values) {
			nodes[i].Left = nodes[left]
		}

		if right < len(values) {
			nodes[i].Right = nodes[right]
		}
	}

	return nodes[0]
}

func PrintPreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PrintPreOrder(root.Left)
	PrintPreOrder(root.Right)
}
