package bruno

import "testing"

func TestCalculateMaximumDepthOfBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected int
	}{
		{
			name: "example case",
			root: &Node{
				value: 3,
				left:  &Node{value: 9},
				right: &Node{
					value: 20,
					left:  &Node{value: 15},
					right: &Node{value: 7},
				},
			},
			expected: 3,
		},
		{
			name:     "single node",
			root:     &Node{value: 1},
			expected: 1,
		},
		{
			name:     "nil root",
			root:     nil,
			expected: 0,
		},
		{
			name: "left-skewed tree",
			root: &Node{
				value: 1,
				left: &Node{
					value: 2,
					left:  &Node{value: 3},
				},
			},
			expected: 3,
		},
		{
			name: "right-skewed tree",
			root: &Node{
				value: 1,
				right: &Node{
					value: 2,
					right: &Node{value: 3},
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateMaximumDepthOfBinaryTree(tt.root)
			if result != tt.expected {
				t.Errorf("calculateMaximumDepthOfBinaryTree() = %d, want %d", result, tt.expected)
			}
		})
	}
}
