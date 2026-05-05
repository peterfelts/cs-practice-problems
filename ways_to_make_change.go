package main

import (
	"fmt"
	"sort"
)

//   Given an array of distinct positive integers representing coin denominations and a
//   single non-negative integer, n, representing a target amount of
//   money, write a function that returns the number of ways to make change for
//   that target amount using the given coin denominations.

//   Example:
//   n = 6
//   denoms = [1, 5]

//   expected output: 2 (1x1 + 1x5, 6x1)

func NumberOfWaysToMakeChange(n int, denoms []int) int {

	// build a slice with length up to the largest denomination
	sort.Ints(denoms)

	fmt.Printf("n: %d, denoms: %d\n", n, denoms)

	if len(denoms) == 0 || n == 0 {
		return 1
	}

	possibleCombinations := make([]int, n+1)

	// set the first value to 1, this represents that there is only
	// one way to make change for a target of zero, which is to return
	// zero coins
	possibleCombinations[0] = 1

	// for each coin we have, we're going to walk through the array of
	// counts of combinations, and apply the formula that  if denomination <= index,
	// then the value at that index is += value @ index (index - denomination)
	for i := 0; i < len(denoms); i++ {
		denom := denoms[i]
		for k := 0; k < len(possibleCombinations); k++ {
			// can we make change for the denomination using the
			// value represented by k (the current index)?
			if denom <= k {
				possibleCombinations[k] += possibleCombinations[k-denom]
			}
		}
		fmt.Println(possibleCombinations)
	}

	fmt.Println(possibleCombinations)
	return possibleCombinations[n]
}
