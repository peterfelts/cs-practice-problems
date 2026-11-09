package binary_search

func binarySearch(arr []int, target int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	idxLeft, idxRight := 0, len(arr)-1

	if idxRight < idxLeft {
		return -1
	}

	for idxLeft <= idxRight {

		// new mid point
		mid := idxLeft + (idxRight-idxLeft)/2

		if arr[mid] == target {
			return mid
		}

		// if the mid point is < target, set left as the mid point
		if arr[mid] < target {
			idxLeft = mid + 1
		} else {
			idxRight = mid - 1
		}
	}

	return -1
}
