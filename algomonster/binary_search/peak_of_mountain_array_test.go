package binary_search

import "testing"

func TestPeakOfMountainArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"example", []int{0, 1, 2, 3, 2, 1, 0}, 3},
		{"three elements", []int{1, 3, 2}, 1},
		{"three elements equal endpoints", []int{0, 1, 0}, 1},
		{"four elements peak near start", []int{1, 5, 3, 2}, 1},
		{"four elements peak near end", []int{1, 3, 5, 2}, 2},
		{"odd length peak at index one", []int{1, 9, 7, 5, 3}, 1},
		{"odd length peak in middle", []int{1, 3, 9, 5, 2}, 2},
		{"odd length peak at penultimate index", []int{1, 3, 5, 9, 2}, 3},
		{"even length peak at index one", []int{1, 10, 8, 6, 4, 2}, 1},
		{"even length peak left of middle", []int{1, 3, 10, 8, 6, 2}, 2},
		{"even length peak right of middle", []int{1, 3, 5, 10, 6, 2}, 3},
		{"even length peak at penultimate index", []int{1, 3, 5, 7, 10, 2}, 4},
		{"long descending slope", []int{0, 10, 9, 8, 7, 6, 5, 4, 3, 2}, 1},
		{"long ascending slope", []int{0, 1, 2, 3, 4, 5, 6, 7, 10, 2}, 8},
		{"all negative values", []int{-10, -5, -1, -3, -8}, 2},
		{"zero peak", []int{-9, -4, 0, -2, -7}, 2},
		{"mixed signs", []int{-8, -2, 5, 12, 3, -1}, 3},
		{"uneven steps", []int{1, 2, 100, 1000, 999, 4, 0}, 3},
		{"same values on opposite slopes", []int{1, 3, 5, 8, 5, 3, 1}, 3},
		{"wide integer range", []int{-2147483648, -1, 2147483647, 0, -2147483647}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := peakOfMountainArray(tt.arr)
			if result != tt.expected {
				t.Errorf("peakOfMountainArray(%v) = %d, want %d", tt.arr, result, tt.expected)
			}
		})
	}
}
