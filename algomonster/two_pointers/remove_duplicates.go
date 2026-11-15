package two_pointers

// Remove Duplicates using Two Pointers
// Given a sorted list of numbers with length at least 1, remove duplicates
// and return the new length. You must do this in-place and without using extra memory.

//                              pointers action
// Input: [0, 0, 1, 1, 1, 2, 2] L=0, R=1 copy arr[1] to arr[0]

// Output: 3.

// Your function should modify the list in place so that the first three elements
// become 0, 1, 2. Return 3 because the new length is 3.

func removeDuplicates(arr []int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	// pattern: two pointers, same direction
	// what is the invariant (the condition that must be true before iterating)?
	// invariant: arr[n-1] != arr[n]

	// use two pointers: write (w) and read (r)
	write, read := 1, 1

	// in case the array is empty or contains only one item
	if len(arr) <= 1 {
		return len(arr)
	}

	for read <= len(arr)-1 {

		// Possible conditions on each iteration:
		// new unique value.
		// this means that arr[read] is unique, so it does
		// belong in the output prefix
		// arr[write-1] != arr[read] || write==0
		// - arr[write] = arr[read]
		// - advance write
		// - advance read
		//
		// duplicate value:
		// this means that arr[read] doesn't belong in the output prefix,
		// so we skip it
		// arr[write-1] == arr[read]
		// - only advance read
		if arr[write-1] != arr[read] {
			arr[write] = arr[read]
			write++
		}
		read++

	}

	return write
}
