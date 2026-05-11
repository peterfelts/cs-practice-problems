package bruno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var findSmallestSubArrayTestCases = []struct {
	paragraph []string
	keywords  []string
	expected  []string
}{
	{
		// second occurrence of first keyword produces a shorter subarray
		paragraph: []string{"how", "green", "now", "green", "brown", "how", "cow", "now", "brown", "cow"},
		keywords:  []string{"how", "now", "brown", "cow"},
		expected:  []string{"how", "cow", "now", "brown", "cow"},
	},
	{
		// keywords appear contiguously inside a longer paragraph
		paragraph: []string{"extra", "a", "b", "c", "extra"},
		keywords:  []string{"a", "b", "c"},
		expected:  []string{"a", "b", "c"},
	},
	{
		// later start yields a shorter match than the first start
		paragraph: []string{"a", "x", "x", "x", "b", "a", "b"},
		keywords:  []string{"a", "b"},
		expected:  []string{"a", "b"},
	},
	{
		// keywords span the entire paragraph with no extra words
		paragraph: []string{"a", "b", "c"},
		keywords:  []string{"a", "b", "c"},
		expected:  []string{"a", "b", "c"},
	},
	{
		// single keyword — returns first occurrence
		paragraph: []string{"a", "b", "a"},
		keywords:  []string{"a"},
		expected:  []string{"a"},
	},
	{
		// shorter subsequence inside a longer subsequence
		paragraph: []string{"a", "j", "k", "b", "a", "b", "c", "z", "d", "a", "c"},
		keywords:  []string{"a", "b", "c"},
		expected:  []string{"a", "b", "c"},
	},
}

func Test_findSmallestSubArrayBruteForce(t *testing.T) {
	for _, tc := range findSmallestSubArrayTestCases {
		actual := findSmallestSubArrayBruteForce(tc.paragraph, tc.keywords)
		assert.Equal(t, tc.expected, actual)
	}
}

func Test_findSmallestSubArrayDP(t *testing.T) {
	for _, tc := range findSmallestSubArrayTestCases {
		actual := findSmallestSubArrayDP(tc.paragraph, tc.keywords)
		assert.Equal(t, tc.expected, actual)
	}
}
