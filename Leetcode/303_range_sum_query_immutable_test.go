package leetcode

import "testing"

func TestNumArraySumRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		left     int
		right    int
		expected int
	}{
		{"example from problem: indices 0-2", []int{-2, 0, 3, -5, 2, -1}, 0, 2, 1},
		{"example from problem: indices 2-5", []int{-2, 0, 3, -5, 2, -1}, 2, 5, -1},
		{"example from problem: indices 0-5", []int{-2, 0, 3, -5, 2, -1}, 0, 5, -3},
		{"single element at start", []int{1, 2, 3}, 0, 0, 1},
		{"single element at end", []int{1, 2, 3}, 2, 2, 3},
		{"entire array of positives", []int{1, 2, 3, 4, 5}, 0, 4, 15},
		{"middle range", []int{1, 2, 3, 4, 5}, 1, 3, 9},
		{"single element array", []int{42}, 0, 0, 42},
		{"all negatives", []int{-3, -1, -4, -1, -5}, 0, 4, -14},
		{"range with zeros", []int{0, 0, 0, 5, 0}, 0, 4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(tt.nums)
			result := obj.SumRange(tt.left, tt.right)
			if result != tt.expected {
				t.Errorf("SumRange(%d, %d) on %v = %d, want %d", tt.left, tt.right, tt.nums, result, tt.expected)
			}
		})
	}
}
