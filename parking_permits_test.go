package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	permits  []int
	expected int
}{
	{
		permits:  []int{1, 3, 2, 2},
		expected: 2,
	},
	{
		permits:  []int{1},
		expected: -1,
	},
	{
		permits:  []int{},
		expected: -1,
	},
	{
		permits:  []int{1, 2, 3, 4, 5, 6, 7, 8, 8},
		expected: 8,
	},
	{
		permits:  []int{1, 2, 3},
		expected: -1,
	},
	{
		permits:  []int{3, 2, 2, 1},
		expected: 2,
	},
	{
		permits:  []int{3, 3, 3, 3},
		expected: 3,
	},
	{
		permits:  []int{3, 1, 2, 4, 9, 6, 7, 6, 5, 8},
		expected: 6,
	},
}

func TestFindDuplicatePermits(t *testing.T) {

	for _, test := range testCases {

		// test the first implementation
		actual := findDuplicatePermits(test.permits)
		require.Equal(t, test.expected, actual)
	}

}

func TestFindDuplicatePermits_ConstantSpace(t *testing.T) {

	for _, test := range testCases {

		// This test sorts the input. Create a copy so that this test
		// won't impact others
		copy := append([]int(nil), test.permits...)
		actual := findDuplicatePermitsConstantSpace(copy)
		require.Equal(t, test.expected, actual)
	}

}

func TestFindDuplicatePermits_ReadOnlyMemory(t *testing.T) {

	for _, test := range testCases {

		testName := fmt.Sprintf("Read-only memory: %v", test.permits)

		t.Run(testName, func(t *testing.T) {
			// test the read-only implementation
			actual := findDuplicatePermitsReadOnlyMemory(test.permits)
			require.Equal(t, test.expected, actual)
		})
	}

}

func TestFindDuplicatePermits_OofNAndOofOneSpace(t *testing.T) {

	for _, test := range testCases {

		testName := fmt.Sprintf("O(1) space and O(n) time: %v", test.permits)
		fmt.Println(testName)

		t.Run(testName, func(t *testing.T) {
			// read only memory, constant space, O(n) time
			actual := findDuplicatePermitsOofNAndOofOneSpace(test.permits)
			require.Equal(t, test.expected, actual)
		})
	}

}
