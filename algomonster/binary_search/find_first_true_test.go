package binary_search

import "testing"

func TestFindBoundary(t *testing.T) {
	tests := []struct {
		name     string
		arr      []bool
		expected int
	}{
		{"boundary in middle", []bool{false, false, true, true, true}, 2},
		{"boundary just after start", []bool{false, true, true, true, true}, 1},
		{"boundary near end", []bool{false, false, false, true, true}, 3},
		{"boundary at end", []bool{false, false, false, false, true}, 4},
		{"all true", []bool{true, true, true, true, true}, 0},
		{"all false", []bool{false, false, false, false, false}, -1},
		{"single true", []bool{true}, 0},
		{"single false", []bool{false}, -1},
		{"two elements boundary at end", []bool{false, true}, 1},
		{"two elements all true", []bool{true, true}, 0},
		{"two elements all false", []bool{false, false}, -1},
		{"even length boundary in middle", []bool{false, false, true, true}, 2},
		{"empty array", []bool{}, -1},
		{"nil array", nil, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFirstTrue(tt.arr)
			if result != tt.expected {
				t.Errorf("findFirstTrue(%v) = %d, want %d", tt.arr, result, tt.expected)
			}
		})
	}
}
