package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"documented geese example", "Do geese see God?", true},
		{"documented car example", "Was it a car or a cat I saw?", true},
		{"documented non-palindrome", "A brown fox jumping over", false},
		{"empty string", "", true},
		{"only non-alphanumeric characters", " \t\n.,!?_-", true},
		{"single letter", "a", true},
		{"single digit", "7", true},
		{"odd-length palindrome", "racecar", true},
		{"even-length palindrome", "abba", true},
		{"mixed case", "AbBa", true},
		{"ignored leading and trailing punctuation", "!racecar?", true},
		{"ignored internal punctuation and whitespace", "a,\tb\n a", true},
		{"numeric palindrome", "12321", true},
		{"letters and digits", "A1b1a", true},
		{"two different letters", "ab", false},
		{"interior mismatch", "abca", false},
		{"numeric non-palindrome", "1231", false},
		{"digits must not be ignored", "0P", false},
		{"punctuation does not hide mismatch", "a, b!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %t, want %t", tt.input, result, tt.expected)
			}
		})
	}
}
