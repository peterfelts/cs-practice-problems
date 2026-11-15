package two_pointers

import (
	"math"
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{"documented example", []int{0, 0, 1, 1, 1, 2, 2}, []int{0, 1, 2}},
		{"single positive element", []int{7}, []int{7}},
		{"single zero", []int{0}, []int{0}},
		{"single negative element", []int{-7}, []int{-7}},
		{"two identical elements", []int{3, 3}, []int{3}},
		{"two distinct elements", []int{1, 2}, []int{1, 2}},
		{"no duplicates", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"all identical", []int{4, 4, 4, 4, 4}, []int{4}},
		{"all zeros", []int{0, 0, 0, 0}, []int{0}},
		{"duplicates at start", []int{1, 1, 1, 2, 3}, []int{1, 2, 3}},
		{"duplicates at end", []int{1, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"duplicates in middle", []int{1, 2, 2, 2, 3}, []int{1, 2, 3}},
		{"equal duplicate groups", []int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}},
		{"unequal duplicate groups", []int{1, 1, 1, 2, 3, 3, 4, 4, 4, 4}, []int{1, 2, 3, 4}},
		{"negative numbers", []int{-5, -5, -3, -3, -1}, []int{-5, -3, -1}},
		{"mixed signs and zero", []int{-3, -3, -1, 0, 0, 2, 2}, []int{-3, -1, 0, 2}},
		{"unique value after long duplicate run", []int{2, 2, 2, 2, 2, 2, 3}, []int{2, 3}},
		{"integer extremes", []int{math.MinInt, math.MinInt, 0, math.MaxInt, math.MaxInt}, []int{math.MinInt, 0, math.MaxInt}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := removeDuplicates(tt.arr)
			if length != len(tt.expected) {
				t.Fatalf("removeDuplicates returned length %d, want %d", length, len(tt.expected))
			}
			if !slices.Equal(tt.arr[:length], tt.expected) {
				t.Errorf("in-place result = %v, want %v", tt.arr[:length], tt.expected)
			}
		})
	}
}
