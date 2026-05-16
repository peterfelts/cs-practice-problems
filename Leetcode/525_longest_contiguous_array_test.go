package leetcode

import "testing"

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"equal halves", []int{0, 1}, 2},
		{"longer equal subarray", []int{0, 1, 0}, 2},
		{"all zeros", []int{0, 0, 0, 0}, 0},
		{"all ones", []int{1, 1, 1, 1}, 0},
		{"entire array is balanced", []int{0, 1, 1, 0}, 4},
		{"unordered but fully balanced", []int{1, 1, 0, 1, 0, 0}, 6},
		{"balanced subarray in the middle", []int{1, 1, 0, 1, 0, 1}, 4},
		{"multiple balanced subarrays, longest wins", []int{0, 1, 0, 0, 1, 1}, 6},
		{"single element zero", []int{0}, 0},
		{"single element one", []int{1}, 0},
		{"empty array", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxLength(tt.nums)
			if result != tt.expected {
				t.Errorf("findMaxLength(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
