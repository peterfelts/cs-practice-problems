package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var removeDuplicatesTestCases = []struct {
	title    string
	nums     []int
	expectedK int
	expectedNums []int
}{
	{
		title:        "empty array",
		nums:         []int{},
		expectedK:    0,
		expectedNums: []int{},
	},
	{
		title:        "single element",
		nums:         []int{7},
		expectedK:    1,
		expectedNums: []int{7},
	},
	{
		title:        "two identical elements",
		nums:         []int{3, 3},
		expectedK:    1,
		expectedNums: []int{3},
	},
	{
		title:        "two distinct elements",
		nums:         []int{1, 2},
		expectedK:    2,
		expectedNums: []int{1, 2},
	},
	{
		title:        "no duplicates",
		nums:         []int{1, 2, 3, 4, 5},
		expectedK:    5,
		expectedNums: []int{1, 2, 3, 4, 5},
	},
	{
		title:        "all duplicates — same value",
		nums:         []int{4, 4, 4, 4},
		expectedK:    1,
		expectedNums: []int{4},
	},
	{
		title:        "duplicates at the start",
		nums:         []int{1, 1, 2, 3},
		expectedK:    3,
		expectedNums: []int{1, 2, 3},
	},
	{
		title:        "duplicates at the end",
		nums:         []int{1, 2, 3, 3},
		expectedK:    3,
		expectedNums: []int{1, 2, 3},
	},
	{
		title:        "duplicates in the middle",
		nums:         []int{1, 2, 2, 3},
		expectedK:    3,
		expectedNums: []int{1, 2, 3},
	},
	{
		title:        "multiple groups of duplicates",
		nums:         []int{1, 1, 2, 2, 3, 3},
		expectedK:    3,
		expectedNums: []int{1, 2, 3},
	},
	{
		title:        "LeetCode example 1",
		nums:         []int{1, 1, 2},
		expectedK:    2,
		expectedNums: []int{1, 2},
	},
	{
		title:        "LeetCode example 2",
		nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		expectedK:    5,
		expectedNums: []int{0, 1, 2, 3, 4},
	},
	{
		title:        "negative numbers with duplicates",
		nums:         []int{-3, -3, -1, 0, 0, 2},
		expectedK:    4,
		expectedNums: []int{-3, -1, 0, 2},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range removeDuplicatesTestCases {
		t.Run(tc.title, func(t *testing.T) {
			k := removeDuplicates(tc.nums)
			assert.Equal(t, tc.expectedK, k)
			assert.Equal(t, tc.expectedNums, tc.nums[:k])
		})
	}
}
