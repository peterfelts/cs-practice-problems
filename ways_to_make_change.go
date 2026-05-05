package main

import (
	"fmt"
	"sort"
)

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
