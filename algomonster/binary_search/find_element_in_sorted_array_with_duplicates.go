package binary_search

// Given a sorted array of integers and a target integer, find the
// first occurrence of the target and return its index.
// Return -1 if the target is not in the array.

// Input:

// arr = [1, 3, 3, 3, 3, 6, 10, 10, 10, 100]
// target = 3
// Output: 1

// Explanation: The first occurrence of 3 is at index 1.

// Input:

// arr = [2, 3, 5, 7, 11, 13, 17, 19]
// target = 6
// Output: -1

// This collapses to a boundary-search problem (FFFTTTT)

func findFirstOccurrence(arr []int, target int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	candidate := -1

	left, right := 0, len(arr)-1

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] >= target {
			right = mid - 1
			candidate = mid
		} else {
			left = mid + 1
		}
	}

	// check if the candiate is valid.
	// If every element is smaller than target, no element is true,
	// and the template returns -1.
	//
	// Thus, one check after the loop is necessary. If first_true_index
	// is not -1 and arr[first_true_index] == target, return first_true_index.
	if candidate > -1 && arr[candidate] == target {
		return candidate
	}

	return -1
}
