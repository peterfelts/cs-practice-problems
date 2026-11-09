package binary_search

import "testing"

func TestFirstNotSmaller(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"first example", []int{1, 3, 3, 5, 8, 8, 10}, 2, 1},
		{"second example", []int{2, 3, 5, 7, 11, 13, 17, 19}, 6, 3},
		{"exact match at start", []int{1, 3, 5, 7}, 1, 0},
		{"exact match in middle", []int{1, 3, 5, 7, 9}, 5, 2},
		{"exact match at end", []int{1, 3, 5, 7}, 7, 3},
		{"target below minimum", []int{1, 3, 5, 7}, 0, 0},
		{"target between last two elements", []int{1, 3, 5, 7}, 6, 3},
		{"first duplicate in middle", []int{1, 3, 3, 3, 5}, 3, 1},
		{"first duplicate at start", []int{2, 2, 2, 4, 6}, 2, 0},
		{"first duplicate near end", []int{1, 3, 5, 5, 5}, 5, 2},
		{"target between duplicate groups", []int{1, 1, 3, 3, 5, 5}, 4, 4},
		{"all equal exact match", []int{3, 3, 3, 3}, 3, 0},
		{"all equal target below", []int{3, 3, 3, 3}, 2, 0},
		{"single element exact match", []int{5}, 5, 0},
		{"single element target below", []int{5}, 4, 0},
		{"two elements exact match at end", []int{2, 4}, 4, 1},
		{"two elements target between", []int{2, 4}, 3, 1},
		{"negative target exact match", []int{-10, -7, -7, -3, -1}, -7, 1},
		{"negative target between elements", []int{-10, -7, -3, -1}, -5, 2},
		{"zero exact match", []int{-3, 0, 0, 2}, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := firstNotSmaller(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("firstNotSmaller(%v, %d) = %d, want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
