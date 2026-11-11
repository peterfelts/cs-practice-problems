package binary_search

import "testing"

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"first example", 16, 4},
		{"second example", 8, 2},
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},
		{"three", 3, 1},
		{"four", 4, 2},
		{"just above four", 5, 2},
		{"nine", 9, 3},
		{"just above nine", 10, 3},
		{"just below sixteen", 15, 3},
		{"just above sixteen", 17, 4},
		{"just below twenty-five", 24, 4},
		{"twenty-five", 25, 5},
		{"just above twenty-five", 26, 5},
		{"just below one hundred", 99, 9},
		{"one hundred", 100, 10},
		{"just above one hundred", 101, 10},
		{"large non-perfect square", 123456789, 11111},
		{"just below large perfect square", 99999999, 9999},
		{"large perfect square", 100000000, 10000},
		{"just above large perfect square", 100000001, 10000},
		{"just below largest 32-bit perfect square", 2147395599, 46339},
		{"largest 32-bit perfect square", 2147395600, 46340},
		{"just above largest 32-bit perfect square", 2147395601, 46340},
		{"maximum signed 32-bit integer", 2147483647, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := squareRoot(tt.n)
			if result != tt.expected {
				t.Errorf("squareRoot(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
