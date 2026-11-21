package two_pointers

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"documented example", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single bar", []int{5}, 0},
		{"two equal bars", []int{5, 5}, 5},
		{"two unequal bars", []int{2, 9}, 2},
		{"all zero heights", []int{0, 0, 0, 0}, 0},
		{"only one nonzero height", []int{0, 5, 0}, 0},
		{"equal heights favor maximum width", []int{4, 4, 4, 4}, 12},
		{"increasing heights", []int{1, 2, 3, 4, 5}, 6},
		{"decreasing heights", []int{5, 4, 3, 2, 1}, 6},
		{"best pair at both ends", []int{5, 1, 1, 1, 5}, 20},
		{"best pair in interior", []int{1, 7, 1, 7, 1}, 14},
		{"zero heights between best pair", []int{0, 5, 0, 5, 0}, 10},
		{"tallest bar is not in best pair", []int{3, 1, 10, 1, 3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containerWithMostWater(tt.arr)
			if result != tt.expected {
				t.Errorf("containerWithMostWater(%v) = %d, want %d", tt.arr, result, tt.expected)
			}
		})
	}
}
