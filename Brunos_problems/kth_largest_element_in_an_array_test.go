package bruno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var kthLargestTestCases = []struct {
	title    string
	values   []int
	k        int
	expected int
}{
	{
		title:    "k=1 returns the largest",
		values:   []int{2, 3, 3, 100, 400, 3, 700, 1},
		k:        1,
		expected: 700,
	},
	{
		title:    "k=4 returns 3 (duplicates count separately)",
		values:   []int{2, 3, 3, 100, 400, 3, 700, 1},
		k:        4,
		expected: 3,
	},
	{
		title:    "k=5 returns 3 (duplicates count separately)",
		values:   []int{2, 3, 3, 100, 400, 3, 700, 1},
		k:        5,
		expected: 3,
	},
	{
		title:    "single element",
		values:   []int{42},
		k:        1,
		expected: 42,
	},
	{
		title:    "two elements, k=1",
		values:   []int{5, 10},
		k:        1,
		expected: 10,
	},
	{
		title:    "two elements, k=2",
		values:   []int{5, 10},
		k:        2,
		expected: 5,
	},
	{
		title:    "all identical elements",
		values:   []int{7, 7, 7, 7},
		k:        3,
		expected: 7,
	},
	{
		title:    "k equals length — returns the smallest",
		values:   []int{3, 1, 4, 1, 5, 9},
		k:        6,
		expected: 1,
	},
	{
		title:    "already sorted ascending",
		values:   []int{1, 2, 3, 4, 5},
		k:        2,
		expected: 4,
	},
	{
		title:    "already sorted descending",
		values:   []int{5, 4, 3, 2, 1},
		k:        3,
		expected: 3,
	},
	{
		title:    "negative numbers",
		values:   []int{-10, -3, -7, -1, -5},
		k:        1,
		expected: -1,
	},
	{
		title:    "mix of negative and positive",
		values:   []int{-5, 0, 3, -2, 8},
		k:        2,
		expected: 3,
	},
}

func TestReturnKthLargestBruteForce(t *testing.T) {
	for _, tc := range kthLargestTestCases {
		t.Run(tc.title, func(t *testing.T) {
			// pass a copy so sort inside the function doesn't affect other tests
			values := make([]int, len(tc.values))
			copy(values, tc.values)
			result := returnKthLargestBruteForce(values, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReturnKthLargestHeap(t *testing.T) {
	for _, tc := range kthLargestTestCases {
		t.Run(tc.title, func(t *testing.T) {
			values := make([]int, len(tc.values))
			copy(values, tc.values)
			result := returnKthLargestHeap(values, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReturnKthLargestQuickSelect(t *testing.T) {
	for _, tc := range kthLargestTestCases {
		t.Run(tc.title, func(t *testing.T) {
			values := make([]int, len(tc.values))
			copy(values, tc.values)
			result := returnKthLargestQuickSelect(values, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}
