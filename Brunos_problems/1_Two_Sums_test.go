package bruno

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSums(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected [2]int
	}{
		{
			name:     "example from problem",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: [2]int{0, 1},
		},
		{
			name:     "answer at end of array",
			nums:     []int{3, 5, 8, 2},
			target:   10,
			expected: [2]int{2, 3},
		},
		{
			name:     "two element array",
			nums:     []int{4, 6},
			target:   10,
			expected: [2]int{0, 1},
		},
		{
			name:     "negative numbers",
			nums:     []int{-3, 4, 7, 1},
			target:   1,
			expected: [2]int{0, 1},
		},
		{
			name:     "target is zero",
			nums:     []int{5, -5, 3, 9},
			target:   0,
			expected: [2]int{0, 1},
		},
		{
			name:     "duplicate values",
			nums:     []int{3, 3, 4, 7},
			target:   6,
			expected: [2]int{0, 1},
		},
		{
			name:     "answer uses first and last element",
			nums:     []int{1, 4, 6, 9},
			target:   10,
			expected: [2]int{0, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSums(tc.nums, tc.target)
			matchesEitherOrder := (result == tc.expected) ||
				(result == [2]int{tc.expected[1], tc.expected[0]})
			require.True(t, matchesEitherOrder,
				"expected indices {%d, %d} in any order, got {%d, %d}",
				tc.expected[0], tc.expected[1], result[0], result[1])
		})
	}
}
