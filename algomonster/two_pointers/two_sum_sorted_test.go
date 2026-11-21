package two_pointers

import (
	"slices"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected []int
	}{
		{"documented example", []int{2, 3, 4, 5, 8, 11, 18}, 8, []int{1, 3}},
		{"two elements", []int{2, 7}, 9, []int{0, 1}},
		{"first and last elements", []int{1, 3, 5, 9}, 10, []int{0, 3}},
		{"first two elements", []int{1, 2, 4, 8}, 3, []int{0, 1}},
		{"last two elements", []int{1, 2, 4, 8}, 12, []int{2, 3}},
		{"adjacent interior elements", []int{1, 3, 4, 8, 12}, 7, []int{1, 2}},
		{"negative numbers", []int{-12, -7, -4, -1}, -11, []int{1, 2}},
		{"mixed signs with zero target", []int{-8, -3, 0, 3, 9}, 0, []int{1, 3}},
		{"pair containing zero", []int{-6, 0, 2, 7}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSumSorted(tt.arr, tt.target)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("twoSumSorted(%v, %d) = %v, want %v", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
