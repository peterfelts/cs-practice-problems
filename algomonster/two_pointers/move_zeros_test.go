package two_pointers

import (
	"slices"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"documented example", []int{1, 0, 2, 0, 0, 7}, []int{1, 2, 7, 0, 0, 0}},
		{"nil slice", nil, nil},
		{"empty slice", []int{}, []int{}},
		{"single zero", []int{0}, []int{0}},
		{"single nonzero", []int{7}, []int{7}},
		{"no zeros", []int{3, 1, 2}, []int{3, 1, 2}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"leading zeros", []int{0, 0, 3, 1}, []int{3, 1, 0, 0}},
		{"trailing zeros", []int{3, 1, 0, 0}, []int{3, 1, 0, 0}},
		{"middle zeros", []int{3, 0, 0, 1}, []int{3, 1, 0, 0}},
		{"alternating zeros", []int{0, 3, 0, 1, 0, 2}, []int{3, 1, 2, 0, 0, 0}},
		{"mixed signs and repeated values", []int{-2, 0, 4, -2, 0, -1}, []int{-2, 4, -2, -1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeros(tt.nums)
			if !slices.Equal(tt.nums, tt.expected) {
				t.Errorf("in-place result = %v, want %v", tt.nums, tt.expected)
			}
		})
	}
}
