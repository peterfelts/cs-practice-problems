package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountOfSubstrings(t *testing.T) {

	// test 1
	// "bcadefa" -> 9
	expected := 9
	actual := countOfSubstrings_Impl("bcadefa", "a")
	require.Equal(t, expected, actual)

	// "a" -> 0
	expected = 0
	actual = countOfSubstrings_Impl("a", "a")
	require.Equal(t, expected, actual)

	// "aaaaaaaa" -> 0
	s := "aaaaaaaa"
	expected = 0
	actual = countOfSubstrings_Impl(s, "a")
	require.Equal(t, expected, actual)

	// "bbbb" -> 10
	s = "bbbb"
	expected = 10
	actual = countOfSubstrings_Impl(s, "a")
	require.Equal(t, expected, actual)

}
