package main

import (
	"fmt"
)

// Write a function that takes in an array of positive integers and
// returns the maximum sum of non-adjacent elements in the array.
//
// array = [75, 105, 120, 75, 90, 135]
// expected = 330

func MaxSubsetSumNoAdjacent(array []int) int {

	fmt.Println(array)

	// special case
	if len(array) == 2 {
		return max(array[0], array[1])
	}
	if len(array) == 3 {
		return max(array[0]+array[2], array[1])
	}

	// how would I brute-force this?
	// I would treat each index like a tree, where the
	// tree would represent the cominations that can be
	// made of non-adjacent values that are TO THE RIGHT
	// of the current index.
	memo := make(map[int]int)
	runningMax := 0
	for i := 0; i < len(array); i++ {
		runningMax = max(runningMax, getMaxSumAtIndex(i, array, memo))
	}

	return runningMax
}

func getMaxSumAtIndex(i int, array []int, memo map[int]int) int {

	// terminal cases
	if i > len(array)-1 {
		return 0
	}
	if i >= len(array)-2 {
		return array[i]
	}

	// memoization
	if storedValue, ok := memo[i]; ok {
		return storedValue
	}

	// calculate the highest value that can be return, for all indices
	runningMax := 0
	for k := i; k < len(array); k++ {
		runningMax = max(runningMax, getMaxSumAtIndex(k+2, array, memo))
	}

	newValue := array[i] + runningMax
	memo[i] = newValue

	return newValue
}
