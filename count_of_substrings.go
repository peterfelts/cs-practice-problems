package main

import (
	"strings"
)

// # Count Substrings Without Letter

// Given a string, `s`, count the number of non-empty substrings of `s` that do not contain the letter 'a'.

// Example: "bcadefa"
// Output: 9. The 9 substrings without 'a' are: "b", "c", "bc", "d", "e", "f", "de", "ef", and "def".

// Constraints:

// - `0 <= s.length <= 10^5`
// - `s` consists of lowercase English letters only

func countOfSubstrings_Impl(s, target string) int {

	parts := strings.Split(s, target)

	countOfSubstrings := 0

	// process each substring.
	// What we're doing here is calculating how many combinations there are
	// for a given substring. For example: 'abc' goes to 'a', 'b', 'c', 'ab', 'bc', and 'abc'
	// the forumual for this i N*(N+1)/2
	// proof: len('abc') = 3 -> 3*(3+1)/2 = 3*4/2 = 12/2 = 6
	for i := 0; i < len(parts); i++ {
		N := len(parts[i])
		countOfSubstrings += N * (N + 1) / 2
	}

	return countOfSubstrings
}
