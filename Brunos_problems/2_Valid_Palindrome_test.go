package bruno

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "example from problem - mixed punctuation and spaces",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "simple palindrome - all lowercase",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "simple palindrome - mixed case",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "even length palindrome",
			input:    "abba",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "not a palindrome - mixed case",
			input:    "Hello",
			expected: false,
		},
		{
			name:     "digits only - palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "digits only - not a palindrome",
			input:    "12345",
			expected: false,
		},
		{
			name:     "alphanumeric with spaces and punctuation",
			input:    "Was it a car or a cat I saw?",
			expected: true,
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidPalindrome(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
