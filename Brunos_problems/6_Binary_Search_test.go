package bruno

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"example case", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"target at start", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"target at end", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"target in middle", []int{-1, 0, 3, 5, 9, 12}, 3, 2},
		{"target not present", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 7, -1},
		{"empty array", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("binarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
