package binary_search

import "testing"

func TestFindMinRotated(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"first example", []int{30, 40, 50, 10, 20}, 3},
		{"second example", []int{3, 5, 7, 11, 13, 17, 19, 2}, 7},
		{"single positive element", []int{5}, 0},
		{"single negative element", []int{-5}, 0},
		{"single zero", []int{0}, 0},
		{"two elements sorted", []int{2, 4}, 0},
		{"two elements rotated", []int{4, 2}, 1},
		{"three elements sorted", []int{1, 3, 5}, 0},
		{"three elements minimum in middle", []int{5, 1, 3}, 1},
		{"three elements minimum at end", []int{3, 5, 1}, 2},
		{"odd length sorted", []int{1, 3, 5, 7, 9}, 0},
		{"odd length minimum at index one", []int{9, 1, 3, 5, 7}, 1},
		{"odd length minimum at index two", []int{7, 9, 1, 3, 5}, 2},
		{"odd length minimum at index three", []int{5, 7, 9, 1, 3}, 3},
		{"odd length minimum at end", []int{3, 5, 7, 9, 1}, 4},
		{"even length sorted", []int{1, 3, 5, 7, 9, 11}, 0},
		{"even length minimum at index one", []int{11, 1, 3, 5, 7, 9}, 1},
		{"even length minimum at index two", []int{9, 11, 1, 3, 5, 7}, 2},
		{"even length minimum at index three", []int{7, 9, 11, 1, 3, 5}, 3},
		{"even length minimum at index four", []int{5, 7, 9, 11, 1, 3}, 4},
		{"even length minimum at end", []int{3, 5, 7, 9, 11, 1}, 5},
		{"all negative sorted", []int{-10, -7, -3, -1}, 0},
		{"all negative rotated", []int{-3, -1, -10, -7}, 2},
		{"mixed signs", []int{2, 8, -9, -4, 0}, 2},
		{"zero is minimum", []int{4, 7, 0, 2}, 2},
		{"widely spaced values", []int{1000000, 2147483647, -2147483648, -100, 0}, 2},
	}

	implementations := []struct {
		name string
		find func([]int) int
	}{
		{"findMinRotated", findMinRotated},
		{"findMinRotatedApproachTwo", findMinRotatedApproachTwo},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result := impl.find(tt.arr)
					if result != tt.expected {
						t.Errorf("%s(%v) = %d, want %d", impl.name, tt.arr, result, tt.expected)
					}
				})
			}
		})
	}
}
