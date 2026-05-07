package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var minNumberOfCoinsForChangeTestCases = []struct {
	n        int
	denoms   []int
	expected int
}{
	{
		// one coin: one 1
		n: 1, denoms: []int{1, 5, 10}, expected: 1,
	},
	{
		// one coin: one 5
		n: 5, denoms: []int{1, 5, 10}, expected: 1,
	},
	{
		// two coins: one 5 + one 1
		n: 6, denoms: []int{1, 5, 10}, expected: 2,
	},
	{
		// one coin: one 10
		n: 10, denoms: []int{1, 5, 10}, expected: 1,
	},
	{
		// two coins: one 10 + one 5
		n: 15, denoms: []int{1, 5, 10}, expected: 2,
	},
	{
		// three coins: one 10 + two 5s (or two 10s + one 5... wait, 25 = 10+10+5)
		n: 25, denoms: []int{5, 10}, expected: 3,
	},
	{
		// impossible — 3 is not reachable with only 2s
		n: 3, denoms: []int{2}, expected: -1,
	},
	{
		// impossible — no denominations provided
		n: 5, denoms: []int{}, expected: -1,
	},
	{
		// zero target — no coins needed
		n: 0, denoms: []int{1, 5, 10}, expected: 0,
	},
	{
		// greedy-trap: naive greedy picks 6+1 (2 coins), but optimal is 4+4 (wait, 11 = 6+5? no)
		// 11 with {1,6,9}: greedy picks 9+1+1=3 coins; optimal is 6+5? no 5 not available
		// actually optimal: 9+1+1=3 or 6+6-? 6+6=12>11, so 9+1+1=3 coins
		n: 11, denoms: []int{1, 6, 9}, expected: 3,
	},
	{
		// greedy-trap: 15 with {1,6,9}: greedy picks 9+6=2 coins; that is optimal
		n: 15, denoms: []int{1, 6, 9}, expected: 2,
	},
}

func TestMinNumberOfCoinsForChange(t *testing.T) {
	for _, tc := range minNumberOfCoinsForChangeTestCases {
		result := MinNumberOfCoinsForChange(tc.n, tc.denoms)
		assert.Equal(t, tc.expected, result)
	}
}
