package binary_search

// An array of boolean values is divided into two sections: The left section
// consists of all false, and the right section consists of all true.
// Find the First True in a Sorted Boolean Array of the right section,
// i.e., the index of the first true element. If there is no true element, return -1.

// Input: arr = [false, false, true, true, true]

// Output: 2

// Explanation: The first true's index is 2.
func findFirstTrue(arr []bool) int {
	// WRITE YOUR BRILLIANT CODE HERE
	left, right := 0, len(arr)-1
	boundary_index := -1

	for left <= right {

		// binary search:
		mid := left + (right-left)/2

		// if the mid-point is true, we disgard all
		// data to the right
		if arr[mid] {
			boundary_index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return boundary_index
}
