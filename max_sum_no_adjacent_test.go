package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxSubsetSumNoAdjacentTestCases = []struct {
	array    []int
	expected int
}{
	{
		// example from problem description: picks indices 0, 2, 5 (75+120+135)
		array: []int{75, 105, 120, 75, 90, 135}, expected: 330,
	},
	{
		// len==2 special case: picks the larger element
		array: []int{1, 2}, expected: 2,
	},
	{
		// len==3: picks first and last (5+5 > 1)
		array: []int{5, 1, 5}, expected: 10,
	},
	{
		// len==3: picking the middle is better than first+last
		array: []int{1, 15, 3}, expected: 15,
	},
	{
		// single element
		array: []int{42}, expected: 42,
	},
	{
		// optimal starts at index 0: picks indices 0, 2, 5 (7+12+14=33)
		array: []int{7, 10, 12, 7, 9, 14}, expected: 33,
	},
	{
		// optimal starts at index 0: picks indices 0, 2, 4, 6 (10+20+15+5=50)
		array: []int{10, 5, 20, 25, 15, 5, 5}, expected: 50,
	},
	{
		// optimal starts at index 0: picks indices 0, 3, 5 (5+100+5=110)
		array: []int{5, 5, 10, 100, 10, 5}, expected: 110,
	},
	{
		// optimal does NOT start at index 0: picks indices 1, 3 (7+6=13 > 3+4+5=12 or 3+6=9)
		array: []int{3, 7, 4, 6, 5}, expected: 13,
	},
	{
		// optimal does NOT start at index 0: picks indices 1, 3 (100+100=200)
		array: []int{1, 100, 1, 100}, expected: 200,
	},
}

func Test_MaxSubsetSumNoAdjacent(t *testing.T) {
	for _, tc := range maxSubsetSumNoAdjacentTestCases {
		actual := MaxSubsetSumNoAdjacent(tc.array)
		assert.Equal(t, tc.expected, actual)
	}
}
