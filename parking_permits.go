package main

import (
	"fmt"
	"sort"
)

/*
Imagine you're managing a parking lot with N parking spaces, each uniquely numbered from 1 to N. Due to a clerical error, you end up issuing N+1 parking permits, each assigned to a parking space number between 1 and N. Since there are more permits than spaces, at least one parking space number must have been assigned to more than one permit. There might even be multiple parking spaces that have duplicate assignments. For example:

If N = 3, the issued permits might list parking spaces [3, 1, 1, 3], meaning both parking spaces 1 and 3 have been assigned more than one permit.
Alternatively, the permits could be [1, 3, 2, 2], indicating that parking space 2 is assigned to multiple permits.

Your task is to identify any one parking space number that has been assigned more than once.

*/

func findDuplicatePermits(inArray []int) int {

	permits := make(map[int]struct{})

	for _, v := range inArray {
		if _, ok := permits[v]; ok { // I was resetting 'v' with my first implementation: if v, ok := permits[v]
			return v // if the key exists, we've found our first duplicate
		} else {
			permits[v] = struct{}{}
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

	// Because the problem specifies the following:
	//   Array size N + 1
	//   Array values 1-N
	//   (N + 1 permits have been issued, but we only have N spaces)
	// We are safe to assume/trust that there will never be a value in
	// the array that is > N. This is important.
	//
	// In this version of the problem, we are constrained in that the array
	// is read-only, and we must deliver a solution that is O(n) time and
	// O(1) space.
	//
	// Again, our goal is to detect duplicates. One way to go about this
	// is to treat the array values as a linked list, conceptually, and
	// to try to detect a cycle. The conceptual linked list will be modeled
	// as each array value being treated as a pointer to another index in the
	// array. So, we're not directly detecting duplicate values, we're indirectly
	// detecting them by the presence or lack of presence of duplicate indicies.
	//
	// We'll use Floyd's Tortise and Hare algorithm to detect a cycle. We'll use
	// two pointers: one fast, one slow. They'll traverse the conceptual linked
	// list. If the two meet, we've found a cycle and then we can determine
	// the value of the duplicate.

	// [3, 1, 3, 4, 2]
	// what does this linked list look like?
	// [0(3), 1(1), 2(3), 3(4), 4(2)]
	// 3 -> 4 -> 2 -> 3 -> 4 -> 2 ...
	slow, fast := 0, 0
	for {

		fmt.Printf("DEBUG:\n\tfast:%d\n\tinArray[%d]:%d\n\n", fast, fast, inArray[fast])

		slow = inArray[slow]
		fast = inArray[inArray[fast]]

		// fmt.Printf("%d(%d)->", slow, inArray[slow])
		// fmt.Printf("%d(%d)->", fast, inArray[fast])

		// we've found our duplicate
		if slow == fast {
			fmt.Println()
			fmt.Printf("slow,fast:%d,%d\n", slow, fast)
			return inArray[slow]
		}

		// we would need to do an input sanitation check to
		// confirm that the array size is N+1 because, otherwise
		// this loop will be infinite.
	}
}
