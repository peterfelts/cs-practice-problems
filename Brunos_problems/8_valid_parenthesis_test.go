package bruno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesValidParentheis = []struct {
	s        string
	expected bool
}{
	{
		s:        "()[]{}",
		expected: true,
	},
	{
		s:        "(]",
		expected: false,
	},
	{
		s:        "([][](){{{[]}}})",
		expected: true,
	},
}

func Test_ValidParenthesis(t *testing.T) {

	for _, test := range testCasesValidParentheis {
		actual := validParenthesis(test.s)
		assert.Equal(t, test.expected, actual)
	}
}
