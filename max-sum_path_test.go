package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var findMaxSumPathTestCases = []struct {
	title    string
	grid     [][]int
	expected int
}{
	{
		title:    "single cell",
		grid:     [][]int{{5}},
		expected: 5,
	},
	{
		title:    "single row",
		grid:     [][]int{{1, 2, 3}},
		expected: 6,
	},
	{
		title:    "single column",
		grid:     [][]int{{1}, {2}, {3}},
		expected: 6,
	},
	{
		// optimal path: 1 -> 4 -> 7 -> 8 -> 9 = 29
		title: "example 1 from problem — prefer going right then down",
		grid: [][]int{
			{1, 4, 3},
			{2, 7, 6},
			{5, 8, 9},
		},
		expected: 29,
	},
	{
		// optimal path: 1 -> 1 -> 1 -> 1 (only one path in a 2x2 of 1s)
		title: "2x2 all ones",
		grid: [][]int{
			{1, 1},
			{1, 1},
		},
		expected: 3,
	},
	{
		// going right: 1->100->1 = 102; going down: 1->1->1 = 3
		title: "2x2 high value in top-right favors going right first",
		grid: [][]int{
			{1, 100},
			{1, 1},
		},
		expected: 102,
	},
	{
		// going down: 1->100->1 = 102; going right: 1->1->1 = 3
		title: "2x2 high value in bottom-left favors going down first",
		grid: [][]int{
			{1, 1},
			{100, 1},
		},
		expected: 102,
	},
	{
		// all paths sum to the same value
		title: "uniform grid — all paths equal",
		grid: [][]int{
			{2, 2, 2},
			{2, 2, 2},
			{2, 2, 2},
		},
		expected: 10,
	},
	{
		// optimal: 1->3->3->3->1 = 11 (go right twice, then down twice)
		title: "large values on top row",
		grid: [][]int{
			{1, 3, 3},
			{1, 1, 3},
			{1, 1, 1},
		},
		expected: 11,
	},
	{
		// single row, single cell with value 1000 (boundary constraint)
		title: "max single-cell value",
		grid:  [][]int{{1000}},
		expected: 1000,
	},
}

func TestFindMaxSumPath(t *testing.T) {
	for _, tc := range findMaxSumPathTestCases {
		t.Run(tc.title, func(t *testing.T) {
			r := len(tc.grid)
			c := len(tc.grid[0])
			memo := make(map[[2]int]int)
			result := findMaxSumPath(tc.grid, 0, 0, r, c, memo)
			assert.Equal(t, tc.expected, result)
		})
	}
}
