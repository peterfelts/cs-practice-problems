package main

import "fmt"

func ZeroSumSubarray(nums []int) bool {
	// walk the array, keeping a running sum.
	// The trick is to understand that if we see
	// the same running sum value at two different points
	// in the array, this must mean that the values between
	// the two incdicies sum to zero:
	// nums: [-2 -3 -1  2  3 4 -5 -3  1  2]
	// sum:   -2 -5 -6 -4 -1 3 -2 -5 -4 -2
	// Here, we see the running sum of -2 three times
	// this means that the sub arrays:
	// [-2] and [-3 -1 2 3 4 -5]
	// sum to zero.
	//
	// The approach will be to keep track of each running sum value
	// If we see a value a second time, we know that we've found
	// two sub arrays tha sum to zero.

	fmt.Println(nums)

	runningSum := 0
	sums := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		runningSum += nums[i]

		if nums[i] == 0 || runningSum == 0 {
			return true
		}

		if _, ok := sums[runningSum]; ok {
			return true
		} else {
			sums[runningSum] = true
		}
	}

	return false
}
