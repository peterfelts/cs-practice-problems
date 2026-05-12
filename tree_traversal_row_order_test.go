package main

import (
	"reflect"
	"testing"
)

var rowOrderTreeTraversalTests = []struct {
	name     string
	tree     *TreeNode
	expected []int
}{
	{
		name:     "nil tree",
		tree:     nil,
		expected: []int{},
	},
	{
		name:     "single node",
		tree:     &TreeNode{value: 1},
		expected: []int{1},
	},
	{
		name: "complete binary tree",
		//       1
		//      / \
		//     2   3
		//    / \ / \
		//   4  5 6  7
		tree: &TreeNode{
			value: 1,
			left: &TreeNode{
				value: 2,
				left:  &TreeNode{value: 4},
				right: &TreeNode{value: 5},
			},
			right: &TreeNode{
				value: 3,
				left:  &TreeNode{value: 6},
				right: &TreeNode{value: 7},
			},
		},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		name: "left-skewed tree",
		//   1
		//  /
		// 2
		//  \
		//   3
		tree: &TreeNode{
			value: 1,
			left: &TreeNode{
				value: 2,
				right: &TreeNode{value: 3},
			},
		},
		expected: []int{1, 2, 3},
	},
	{
		name: "right-skewed tree",
		// 1
		//  \
		//   2
		//    \
		//     3
		tree: &TreeNode{
			value: 1,
			right: &TreeNode{
				value: 2,
				right: &TreeNode{value: 3},
			},
		},
		expected: []int{1, 2, 3},
	},
	{
		name: "unbalanced tree",
		//       10
		//      /  \
		//     5    20
		//    /
		//   3
		//  /
		// 1
		tree: &TreeNode{
			value: 10,
			left: &TreeNode{
				value: 5,
				left: &TreeNode{
					value: 3,
					left:  &TreeNode{value: 1},
				},
			},
			right: &TreeNode{value: 20},
		},
		expected: []int{10, 5, 20, 3, 1},
	},
}

func TestRowOrderTreeTraversal(t *testing.T) {
	for _, tt := range rowOrderTreeTraversalTests {
		t.Run(tt.name, func(t *testing.T) {
			got := rowOrderTreeTraversal(tt.tree)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rowOrderTreeTraversal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRowOrderTreeTraversalSlice(t *testing.T) {
	for _, tt := range rowOrderTreeTraversalTests {
		t.Run(tt.name, func(t *testing.T) {
			got := rowOrderTreeTraversalSlice(tt.tree)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rowOrderTreeTraversalSlice() = %v, want %v", got, tt.expected)
			}
		})
	}
}
