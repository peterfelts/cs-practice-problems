package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_To_Int_Base10(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		panics   bool
	}{
		{"basic positive", "123", 123, false},
		{"single digit", "7", 7, false},
		{"zero", "0", 0, false},
		{"negative number", "-456", -456, false},
		{"negative single digit", "-9", -9, false},
		{"large number", "2147483647", 2147483647, false},
		{"large negative", "-2147483648", -2147483648, false},
		{"one", "1", 1, false},
		{"ten", "10", 10, false},
		{"one hundred", "100", 100, false},
		{"empty string panics", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { stoi(tt.input, MODE_BASE10) })
			} else {
				assert.Equal(t, tt.expected, stoi(tt.input, MODE_BASE10))
			}
		})
	}
}

func Test_String_To_Int_Hex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		panics   bool
	}{
		{"single digit", "A", 10, false},
		{"single digit B", "B", 11, false},
		{"single digit F", "F", 15, false},
		{"single digit zero", "0", 0, false},
		{"single digit numeric", "9", 9, false},
		{"two digits all hex", "FF", 255, false},
		{"two digits mixed", "1A", 26, false},
		{"three digits", "1FF", 511, false},
		{"four digits", "FFFF", 65535, false},
		{"negative hex", "-1A", -26, false},
		{"empty string panics", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { stoi(tt.input, MODE_HEX) })
			} else {
				assert.Equal(t, tt.expected, stoi(tt.input, MODE_HEX))
			}
		})
	}
}
