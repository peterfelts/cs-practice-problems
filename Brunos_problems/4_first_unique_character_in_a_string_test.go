package bruno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem: Given a string s, find the first non-repeating character in it and return its index.
// If it does not exist, return -1.

// Examples:
// Input: s = "leetcode"
// Output: 0
// Input: s = "loveleetcode"
// Output: 2

var testCasesFirstNonRepeating = []struct {
	s        string
	expected int
}{
	{
		s:        "leetcode",
		expected: 0,
	},
	{
		s:        "loveleetcode",
		expected: 2,
	},
	{
		s:        "aabbccddeeffgg",
		expected: -1,
	},
}

func Test_firstNonRepeatingCharacter(t *testing.T) {

	for _, testCase := range testCasesFirstNonRepeating {
		actual := firstNonRepeatingCharacter(testCase.s)

		assert.Equal(t, testCase.expected, actual)
	}
}
