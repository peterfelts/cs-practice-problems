package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"target in middle", []int{-1, 0, 3, 5, 9, 12}, 3, 2},
		{"target at start", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"target at end", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"target between elements", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"target below minimum", []int{-1, 0, 3, 5, 9, 12}, -2, -1},
		{"target above maximum", []int{-1, 0, 3, 5, 9, 12}, 13, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element target below", []int{5}, 4, -1},
		{"single element target above", []int{5}, 6, -1},
		{"two elements target at start", []int{2, 4}, 2, 0},
		{"two elements target at end", []int{2, 4}, 4, 1},
		{"two elements target missing", []int{2, 4}, 3, -1},
		{"negative numbers", []int{-10, -7, -3, -1}, -7, 1},
		{"empty array", []int{}, 1, -1},
		{"nil array", nil, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("binarySearch(%v, %d) = %d, want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
