package leetcode

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{"example 1: PAYPALISHIRING 3 rows", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"example 2: PAYPALISHIRING 4 rows", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"single row", "ABCDE", 1, "ABCDE"},
		{"two rows", "ABCDE", 2, "ACEBD"},
		{"numRows >= string length", "AB", 3, "AB"},
		{"single character", "A", 1, "A"},
		{"single character more rows", "A", 5, "A"},
		{"numRows equals string length", "ABCD", 4, "ABCD"},
		{"all same characters", "AAAA", 2, "AAAA"},
		{"two rows longer string", "ABCDEFGH", 2, "ACEGBDFH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convert(tt.s, tt.numRows)
			if result != tt.expected {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, result, tt.expected)
			}
		})
	}
}
