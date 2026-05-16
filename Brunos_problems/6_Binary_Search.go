package bruno

import "fmt"

// 6. Binary Search (Algorithms)
// Problem: Given a sorted (in ascending order) integer array nums of n elements and a target value,
// write a function to search for target in nums. If target exists, then return its index. Otherwise, return -1.

// Example:

// Input: nums = [-1,0,3,5,9,12], target = 9

// Output: 4

const NOT_FOUND int = -1

func binarySearch(nums []int, target int) int {

	if len(nums) == 0 {
		return NOT_FOUND
	}

	l := 0
	r := len(nums) - 1

	for {

		if l == r {
			if nums[l] == target {
				return l
			}
			return NOT_FOUND
		}

		pivot := l + (r-l)/2
		fmt.Printf("l, r: (%d, %d), pivot: %d\n", l, r, pivot)

		if nums[pivot] == target {
			return pivot

			// search right
		} else if nums[pivot] > target {
			r = pivot
			// search right
		} else {
			l = pivot + 1

		}

	}
}
