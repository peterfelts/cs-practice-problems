package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var waysToMakeChangeTestCases = []struct {
	n        int
	denoms   []int
	expected int
}{
	{
		// two ways: six 1s, or one 5 + one 1
		n: 6, denoms: []int{1, 5}, expected: 2,
	},
	{
		// target is zero — exactly one way (use no coins)
		n: 0, denoms: []int{2, 3, 4, 7}, expected: 1,
	},
	{
		// 9 is not divisible by 5 — no way to make change
		n: 9, denoms: []int{5}, expected: 0,
	},
	{
		// odd target with only even coins — impossible
		n: 7, denoms: []int{2, 4}, expected: 0,
	},
	{
		// only one way: four 1s (5, 10, 25 all exceed target)
		n: 4, denoms: []int{1, 5, 10, 25}, expected: 1,
	},
	{
		// two ways: five 1s, or one 5
		n: 5, denoms: []int{1, 5, 10, 25}, expected: 2,
	},
	{
		// four ways: ten 1s | 5+five 1s | 5+5 | 10
		n: 10, denoms: []int{1, 5, 10, 25}, expected: 4,
	},
	{
		// thirteen ways to make 25 cents with standard US coins
		n: 25, denoms: []int{1, 5, 10, 25}, expected: 13,
	},
	{
		// four ways: 2+2+2+2+2+2 | 2+2+2+3+3 | 3+3+3+3 | 2+3+7
		n: 12, denoms: []int{2, 3, 7}, expected: 4,
	},
	{
		// three ways: 7 | 3+4 | 2+2+3
		n: 7, denoms: []int{2, 3, 4, 7}, expected: 3,
	},
}

func Test_NumberOfWaysToMakeChange(t *testing.T) {
	for _, tc := range waysToMakeChangeTestCases {
		actual := NumberOfWaysToMakeChange(tc.n, tc.denoms)
		assert.Equal(t, tc.expected, actual)
	}
}
