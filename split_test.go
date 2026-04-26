package main

import (
	"fmt"
	"slices"
	"testing"
)

type splitTest struct {
	input     string
	delimeter string
	expected  []string
}

var splitTestCases = []splitTest{
	{
		input:     "split by space",
		delimeter: " ",
		expected:  []string{"split", "by", "space"},
	},
	{
		input:     "beekeeper needed",
		delimeter: "e",
		expected:  []string{"b", "", "k", "", "p", "r n", "", "d", "d"},
	},
	{
		input:     "/home/./..//Documents/",
		delimeter: "/",
		expected:  []string{"", "home", ".", "..", "", "Documents", ""},
	},
	{
		input:     "",
		delimeter: "?",
		expected:  []string{},
	},
	{
		input:     "abcdefg",
		delimeter: "?",
		expected:  []string{"abcdefg"},
	},
	{
		input:     "",
		delimeter: "",
		expected:  []string{},
	},
	{
		input:     "a",
		delimeter: "a",
		expected:  []string{"", ""},
	},
	{
		input:     "abcdefg",
		delimeter: "",
		expected:  []string{"abcdefg"},
	},
}

func Test_Split(t *testing.T) {

	for _, test := range splitTestCases {
		actual := split(test.input, test.delimeter)
		if !slices.Equal(test.expected, actual) {
			fmt.Printf("FAIL: expected: %v, actaul: %v\n", test.expected, actual)
			fmt.Printf("len(expected): %d, len(actual): %d\n", len(test.expected), len(actual))
		} else {
			fmt.Printf("PASS: expected: %s\n", test.expected)
		}
	}
}
