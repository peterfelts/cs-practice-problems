package main

import (
	"sort"
)

/*
Imagine you're managing a parking lot with N parking spaces, each uniquely numbered from 1 to N. Due to a clerical error, you end up issuing N+1 parking permits, each assigned to a parking space number between 1 and N. Since there are more permits than spaces, at least one parking space number must have been assigned to more than one permit. There might even be multiple parking spaces that have duplicate assignments. For example:

If N = 3, the issued permits might list parking spaces [3, 1, 1, 3], meaning both parking spaces 1 and 3 have been assigned more than one permit.
Alternatively, the permits could be [1, 3, 2, 2], indicating that parking space 2 is assigned to multiple permits.

Your task is to identify any one parking space number that has been assigned more than once.

*/

func findDuplicatePermits(inArray []int) int {

	permits := make(map[int]int)

	for _, v := range inArray {
		if _, ok := permits[v]; ok { // I was resetting 'v' with my first implementation: if v, ok := permits[v]
			return v // if the key exists, we've found our first duplicate
		} else {
			permits[v] = 1
		}
	}

	return -1

}

/*
1 - Fix the solution above.
2 - What if we do not have enough memory to create a copy of the input (can not be O(n) space).
3 - Same as 2 but input is in rad-only memory (can not modify inArray).
4 - Finally, same as 2 and 3 but we must do it in O(n) time (and O(1) space).
*/

// #2
// constant space
func findDuplicatePermitsConstantSpace(inArray []int) int {

	// Sort the array, walk it, and return the first duplicate value I find
	sort.Ints(inArray)

	// store the previous value
	prev := -1
	for i := 0; i < len(inArray); i++ {
		if prev == inArray[i] {
			return prev
		}
		prev = inArray[i]
	}

	return -1

}

// #3
// read-only memory.
// This implementation will be comupte heavy as we are unable to
// copy the input array and we are also constained in that we may not
// modify the array
func findDuplicatePermitsReadOnlyMemory(inArray []int) int {

	for i := 0; i < len(inArray); i++ {
		for j := 0; j < len(inArray); j++ {
			if inArray[i] == inArray[j] && i != j {
				return inArray[i]
			}
		}
	}

	return -1
}

// #4
// inArray is read only, solution must be: O(n) time (and O(1) space)
func findDuplicatePermitsOofNAndOofOneSpace(inArray []int) int {

	// I was unable to figure this out unless we know that the array is given
	// to us sorted. I would ask the interviewer this question.

	// If I could count on inArray being sorted and starting at 1, I could walk the array
	// and used the closed form k*(k+1)/2 on each step, keeping a manual sum.
	// when the manual sume was greater than the closed form (k*(k+1)/2), I would have
	// identified my first duplicate. On the other hand, this isn't much better than the
	// implementation for #2
	runningSum := 0
	for i := 0; i < len(inArray); i++ {
		runningSum += inArray[i]
		k := i + 1
		closedForm := k * (k + 1) / 2
		if runningSum != closedForm {
			return inArray[i]
		}
	}

	return -1
}
