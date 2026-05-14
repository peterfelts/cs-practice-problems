package interviewingio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// buildExampleTree builds the tree from the problem description:
//
//	      O
//	     /
//	    O
//	   / \
//	  O   O
//	 / \   \
//	O   O   O
func buildExampleTree() *Node {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 4}
	root.Left.Left.Left = &Node{Value: 6}
	root.Left.Left.Right = &Node{Value: 7}
	root.Left.Right.Right = &Node{Value: 5}
	return root
}

func TestCalculateProlificness_NilTree(t *testing.T) {
	assert.Equal(t, -1, calculateProlificness(nil))
}

func TestCalculateProlificness_ExampleTree(t *testing.T) {
	// Level 0: 1 node, 1 child  → prolificness 1.0
	// Level 1: 1 node, 2 children → prolificness 2.0  ← most prolific
	// Level 2: 2 nodes, 3 children → prolificness 1.5
	// Level 3: 3 nodes, 0 children → prolificness 0.0
	assert.Equal(t, 1, calculateProlificness(buildExampleTree()))
}

func TestCalculateProlificness_SingleNode(t *testing.T) {
	// Only level 0 exists; no children, so prolificness cannot be computed.
	// Returns default maxLevel 0.
	root := &Node{Value: 1}
	assert.Equal(t, 0, calculateProlificness(root))
}

func TestCalculateProlificness_RootWithTwoChildren(t *testing.T) {
	// Level 0: 1 node, 2 children → prolificness 2.0  ← most prolific
	// Level 1: 2 nodes, 0 children → prolificness 0.0
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	assert.Equal(t, 0, calculateProlificness(root))
}

func TestCalculateProlificness_RightSkewed(t *testing.T) {
	// root → right → right
	// Level 0: 1 node, 1 child → prolificness 1.0
	// Level 1: 1 node, 1 child → prolificness 1.0
	// Level 2: 1 node, 0 children → prolificness 0.0
	// Tie between levels 0 and 1; level 0 wins (first found).
	root := &Node{Value: 1}
	root.Right = &Node{Value: 2}
	root.Right.Right = &Node{Value: 3}
	assert.Equal(t, 0, calculateProlificness(root))
}

func TestCalculateProlificness_CompleteTwoLevels(t *testing.T) {
	// Perfect binary tree with 3 levels (root + 2 + 4 nodes).
	// Level 0: 1 node, 2 children → prolificness 2.0
	// Level 1: 2 nodes, 4 children → prolificness 2.0
	// Level 2: 4 nodes, 0 children → prolificness 0.0
	// Tie between levels 0 and 1; level 0 wins (strict > keeps first).
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}
	assert.Equal(t, 0, calculateProlificness(root))
}
