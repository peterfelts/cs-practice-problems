package binary_search

import "testing"

func TestFindFirstOccurrence(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"first example", []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}, 3, 1},
		{"second example", []int{2, 3, 5, 7, 11, 13, 17, 19}, 6, -1},
		{"unique target at start", []int{1, 3, 5, 7, 9}, 1, 0},
		{"unique target in middle", []int{1, 3, 5, 7, 9}, 5, 2},
		{"unique target at end", []int{1, 3, 5, 7, 9}, 9, 4},
		{"duplicates at start", []int{2, 2, 2, 4, 6}, 2, 0},
		{"duplicates in middle", []int{1, 3, 3, 3, 5}, 3, 1},
		{"duplicates at end", []int{1, 3, 5, 5, 5}, 5, 2},
		{"even length duplicate group", []int{1, 2, 2, 2, 2, 3}, 2, 1},
		{"target between duplicate groups", []int{1, 1, 3, 3, 5, 5}, 4, -1},
		{"target below minimum", []int{1, 3, 3, 5}, 0, -1},
		{"target above maximum", []int{1, 3, 3, 5}, 6, -1},
		{"all equal target found", []int{3, 3, 3, 3}, 3, 0},
		{"all equal target below", []int{3, 3, 3, 3}, 2, -1},
		{"all equal target above", []int{3, 3, 3, 3}, 4, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element target below", []int{5}, 4, -1},
		{"single element target above", []int{5}, 6, -1},
		{"two elements target at start", []int{2, 4}, 2, 0},
		{"two elements target at end", []int{2, 4}, 4, 1},
		{"two elements target missing", []int{2, 4}, 3, -1},
		{"two equal elements", []int{2, 2}, 2, 0},
		{"negative target found", []int{-10, -7, -7, -3, -1}, -7, 1},
		{"negative target missing", []int{-10, -7, -3, -1}, -5, -1},
		{"zero target found", []int{-3, 0, 0, 2}, 0, 1},
		{"empty array", []int{}, 1, -1},
		{"nil array", nil, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFirstOccurrence(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("findFirstOccurrence(%v, %d) = %d, want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
