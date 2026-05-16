package bruno

import "testing"

func TestContainsDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"example case - has duplicate", []int{1, 2, 3, 1}, true},
		{"no duplicates", []int{1, 2, 3, 4}, false},
		{"all duplicates", []int{1, 1, 1, 1}, true},
		{"duplicate at end", []int{1, 2, 3, 4, 2}, true},
		{"single element", []int{1}, false},
		{"empty array", []int{}, false},
		{"negative numbers with duplicate", []int{-1, -2, -3, -1}, true},
		{"negative numbers no duplicate", []int{-1, -2, -3, -4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicates(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicates(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
