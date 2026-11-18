package two_pointers

// Move Zeros
// Given an array of integers, move all the 0s to the back of the array while
// maintaining the relative order of the non-zero elements. Do this in-place using
// constant auxiliary space.

// Input:
// [1, 0, 2, 0, 0, 7]
//  1, 2, 0, 0, 0, 7
//  1, 2,

// Output:
// [1, 2, 7, 0, 0, 0]

// [0, 0, 0, 0, 0, 1]
//  0  0  0  0  1  0
//  0  0  0  1  0  0

func moveZeros(nums []int) {

	slow, fast := 0, 0

	// Nothing to do.
	if len(nums) < 2 {
		return
	}

	// Hold the slow pointer when it is on a zero.
	// Only swap the values at slow and fast when fast
	// points at a non-zero value. At that point,
	// swap the items and increment slow.
	for fast < len(nums) {

		if nums[slow] == 0 {
			if nums[fast] != 0 {
				// swap
				tmp := nums[slow]
				nums[slow] = nums[fast]
				nums[fast] = tmp
				slow++
			}
		} else {
			slow++
		}

		fast++
	}
}
