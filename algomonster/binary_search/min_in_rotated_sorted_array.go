package binary_search

// A sorted array of unique integers was rotated at an unknown pivot.
// For example, [10, 20, 30, 40, 50] becomes [30, 40, 50, 10, 20].
// Find the index of the minimum element in this array.

// Input: [30, 40, 50, 10, 20]

// Output: 3

// Explanation: The smallest element is 10, and its index is 3.

// Input: [3, 5, 7, 11, 13, 17, 19, 2]

// Output: 7

// Explanation: The smallest element is 2, and its index is 7.

func findMinRotated(arr []int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	// binary search but instead of only looking at min,
	// the feasibility function will be comparing the L and R
	// values, to determine if we've at least found a range
	// where the boundary exists
	left, right := 0, len(arr)-1

	for left < right {

		m := left + (right-left)/2

		if arr[m] > arr[m+1] {
			return m + 1
		}

		// the boundary is to the right
		if arr[m] > arr[right] {
			left = m + 1

		} else {
			// we don't reduce the right range
			// as we are doing comparisons to find the
			// boundary (we need to consider the right in our
			// next iteration)
			right = m
		}
	}

	if left == right {
		return left
	}

	return 0
}

// Another approach is to compare each value to the last value in the input array.
// This then reduces to a boundary-search problem:
//
// <= 20
//
// [30 40 50 10 20]
// [ F  F  F  T  T]
func findMinRotatedApproachTwo(arr []int) int {

	left, right := 0, len(arr)-1
	candidate := 0

	for left <= right {

		m := left + (right-left)/2

		// is the value at the mid point <= the value at the end of the input array.
		// If so, look to the left, otherwise, look right
		if arr[m] <= arr[len(arr)-1] {
			candidate = m
			right = m - 1
		} else {
			left = m + 1
		}

	}

	return candidate
}
