package bruno

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var productTestCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "no zeros - example from problem",
		input:    []int{1, 2, 3, 4},
		expected: []int{24, 12, 8, 6},
	},
	{
		name:     "no zeros - all ones",
		input:    []int{1, 1, 1, 1},
		expected: []int{1, 1, 1, 1},
	},
	{
		name:     "no zeros - two elements",
		input:    []int{3, 7},
		expected: []int{7, 3},
	},
	{
		name:     "no zeros - negative values",
		input:    []int{-1, 2, -3, 4},
		expected: []int{-24, 12, -8, 6},
	},
	{
		name:     "one zero - zero at start",
		input:    []int{0, 2, 3, 4},
		expected: []int{24, 0, 0, 0},
	},
	{
		name:     "one zero - zero in middle",
		input:    []int{1, 2, 0, 4},
		expected: []int{0, 0, 8, 0},
	},
	{
		name:     "one zero - zero at end",
		input:    []int{1, 2, 3, 0},
		expected: []int{0, 0, 0, 6},
	},
	{
		name:     "multiple zeros - two zeros",
		input:    []int{0, 2, 0, 4},
		expected: []int{0, 0, 0, 0},
	},
	{
		name:     "multiple zeros - all zeros",
		input:    []int{0, 0, 0},
		expected: []int{0, 0, 0},
	},
	{
		name:     "multiple zeros - three zeros among values",
		input:    []int{5, 0, 3, 0, 2, 0},
		expected: []int{0, 0, 0, 0, 0, 0},
	},
}

func TestProductOfArrayUsingDivision(t *testing.T) {
	for _, tc := range productTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := productOfArrayUsingDivision(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestProductOfArrayNoDivision(t *testing.T) {
	for _, tc := range productTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := productOfArrayNoDivision(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
