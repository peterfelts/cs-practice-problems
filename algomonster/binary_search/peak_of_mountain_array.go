package binary_search

// A mountain array is defined as an array that

// has at least 3 elements
// has an element with the largest value called "peak", with index k. The
// array elements strictly increase from the first element to A[k], and then
// strictly decrease from A[k + 1] to the last element of the array. Thus
// creating a "mountain" of numbers.
// That is, given A[0]<...<A[k-1]<A[k]>A[k+1]>...>A[n-1], we need to find
// the index k. Note that the peak element is neither the first nor the
// lastIndex of the array.

// Find the index of the peak element. Assume there is only one peak element.

// Input: 0 1 2 3 2 1 0

// Output: 3

// Explanation: The largest element is 3, and its index is 3.

func peakOfMountainArray(arr []int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	left, right := 0, len(arr)-1
	candidate := 0

	for left <= right {
		m := left + (right-left)/2

		// If the value at the mid point > the value of the next item,
		// then we know that we're somewhere on the right-hand-side
		// of the mountain, so we need to keep looking to the left
		// until we find the peak
		if arr[m] > arr[m+1] {
			candidate = m
			right = m - 1
		} else {
			left = m + 1
		}

	}

	return candidate
}
