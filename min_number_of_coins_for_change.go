package main

import (
	"sort"
)

// Given an array of positive integers representing coin denominations and a
// single non-negative integer n representing a target amount of
// money, write a function that returns the smallest number of coins needed to
// make change for (to sum up to) that target amount using the given coin
// denominations.

// Note that you have access to an unlimited amount of coins. In other words, if
// the denominations are [1, 5, 10], you have access to an unlimited
// amount of 1s, 5s, and 10s.

// If it's impossible to make change for the target amount, return -1.

const MAX_UINT uint = ^uint(0)

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Create an array that is length n + 1
	// Populate array with increasing linear values 1, 2, 3...
	// These represent valuations
	// Iterate over all denominations. For each denomination,
	// iterate across the entire list (L).
	// the formula is:
	//
	// if( current denomination < i ) {
	//     L[i] = min(L[i], 1 + L[i - current denomination])
	// }
	//
	// Why 1 + L[i - current denomination]?
	// We're saying that with the current denomination, we can
	// achieve the current value (i) by using one coin at the
	// current denomination, but we'll be left with a remainder.
	// How do we solve for the remainder? Well, the good news is
	// that we've already solved for it!
	// For example, let's say we're trying to reach a value for 4
	// and we're currently iterating over the array with a denomination
	// of 2. 4 - one coin of value 2 = 4 - 2 = 2. We'll have already
	// solved for how to make a value of 2 with a denomination of 2
	// so we just look that up

	sort.Ints(denoms)

	coinCounts := make([]uint, n+1)
	for i := 0; i < len(coinCounts); i++ {
		coinCounts[i] = MAX_UINT
	}
	coinCounts[0] = 0

	// iterate over denominations
	for _, d := range denoms {

		// iterate over all possible values: 0 ... n
		for amount, _ := range coinCounts {

			// if the current denomination is less than the
			// current value,
			if d <= amount {

				// It looks like, in Go, when I add 1 + MAX_UINT, this causes a rollover
				// without an exception/error. This will break our algorithm as we'll see
				// unexpected values in some parts of the array, while we're using MAX_UINT
				// to indicate an unset value. To account for this, we will only set the
				// potential new value to 1 + coinCounts[amount -d] when the [amount-d]
				// value IS NOT MAX_UINT
				prevCoinCount := 1 + coinCounts[amount-d]
				if coinCounts[amount-d] == MAX_UINT {
					prevCoinCount = MAX_UINT
				}

				coinCounts[amount] = minUint(coinCounts[amount], prevCoinCount)
			}
		}
	}

	if coinCounts[n] == MAX_UINT {
		return -1
	}

	return (int)(coinCounts[n])
}

func minUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
