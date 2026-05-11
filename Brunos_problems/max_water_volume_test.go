package bruno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxWaterVolumeBruteForceTestCases = []struct {
	title    string
	heights  []int
	expected int
}{
	{
		title:   "empty — no water",
		heights: []int{}, expected: 0,
	},
	{
		title:   "single bar — no water",
		heights: []int{5}, expected: 0,
	},
	{
		title:   "two bars — no water (open on both sides)",
		heights: []int{3, 4}, expected: 0,
	},
	{
		title:   "ascending — no walls on the right to trap water",
		heights: []int{1, 2, 3, 4, 5}, expected: 0,
	},
	{
		title:   "descending — no walls on the left to trap water",
		heights: []int{5, 4, 3, 2, 1}, expected: 0,
	},
	{
		title:   "simple valley: 1 unit trapped at the bottom",
		heights: []int{1, 0, 1}, expected: 1,
	},
	{
		title:   "symmetric valley: 3 units trapped at center",
		heights: []int{3, 0, 3}, expected: 3,
	},
	{
		// index 2 traps 1 unit (between h=1 and h=5),
		// index 4 traps 4 units (between h=5 and h=5, bar height is 1)
		// total = 5 (note: the source comment incorrectly states 4)
		title:   "problem comment example",
		heights: []int{0, 1, 0, 5, 1, 5}, expected: 5,
	},
	{
		title:   "classic LeetCode #42 example",
		heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6,
	},
	{
		title:   "multiple valleys at varying depths",
		heights: []int{3, 0, 2, 0, 4}, expected: 7,
	},
	{
		title:   "tall left wall, sloped right side",
		heights: []int{4, 2, 0, 3, 2, 5}, expected: 9,
	},
	{
		title:   "flat plateau — no water",
		heights: []int{3, 3, 3, 3}, expected: 0,
	},
	{
		title:   "single dip between equal walls",
		heights: []int{2, 0, 2}, expected: 2,
	},
}

func TestMaxWaterVolumeBruteForce(t *testing.T) {
	for _, tc := range maxWaterVolumeBruteForceTestCases {
		t.Run(tc.title, func(t *testing.T) {
			result := maxWaterVolumeBruteForce(tc.heights)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMaxWaterVolumeBruteForceDP(t *testing.T) {
	for _, tc := range maxWaterVolumeBruteForceTestCases {
		t.Run(tc.title, func(t *testing.T) {
			result := maxWaterVolumeDP(tc.heights)
			assert.Equal(t, tc.expected, result)
		})
	}
}
