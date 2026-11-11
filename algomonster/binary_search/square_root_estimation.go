package binary_search

// Given an integer, find its square root without using the built-in
// square root function. Only return the integer part (truncate the decimals).

// Input: 16

// Output: 4

// Input: 8

// Output: 2

// Explanation: square root of 8 is 2.83..., return the integer part, 2

func squareRoot(n int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	left, right := 0, n
	candidate := n

	// store the delta between n and the approximation
	delta := n

	for left <= right {
		mid := left + (right-left)/2

		// if the square of the mid is > than n, look left
		mid_square := mid * mid
		if mid_square >= n {
			right = mid - 1
		} else {
			left = mid + 1
		}

		// update the detla if it gets smaller
		delta_tmp := n - mid_square
		if delta_tmp >= 0 && delta_tmp < delta {
			delta = delta_tmp
			candidate = mid
		}
	}

	return candidate
}
