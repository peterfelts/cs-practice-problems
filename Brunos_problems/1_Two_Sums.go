package bruno

import (
	"sort"
)

// Problem: Given an array of integers nums and an integer target, return the indices of the two
// numbers such that they add up to target. You may assume that each input would have exactly one
// solution, and you may not use the same element twice.

// Example:

// Input: nums = [2, 7, 11, 15], target = 9

// Output: [0, 1] (Since $nums[0] + nums[1] = 2 + 7 = 9$)

type item struct {
	value         int
	originalindex int
}

// NOTE: my solution returns the indices in sorted order, which is not
// necessarily the original order that the indices would be in. The
// problem didn't specify that the indices must be returned in the
// original order.
func twoSums(nums []int, target int) [2]int {

	out := [2]int{-1, -1}

	// The question asks for the indices that sum to the target,
	// not the values that sum to the target. Searching for all
	// permutations would be O(n^2), and I don't want to do that,
	// so what I will do instead is to store the original index
	// for each value in a struct, then create an array of structs,
	// then sort that. Once we've found our targets, we can simply
	// return the corresponding indices in the original order.
	numsWithIndex := make([]item, len(nums))
	for i := 0; i < len(nums); i++ {
		numsWithIndex[i] = item{nums[i], i}
	}

	sort.Slice(numsWithIndex, func(a, b int) bool {
		return numsWithIndex[a].value < numsWithIndex[b].value
	})

	// use two pointers, a left and right
	// start left at index 0, right at the last
	// index.
	// Sum the values at the two pointers.
	// If the sum is > target, decrement right.
	// If the sum is < target, increment left
	// break when we find the target or when the two
	// indices meet
	leftPointer := 0
	rightPointer := len(nums) - 1
	for {

		sum := numsWithIndex[leftPointer].value + numsWithIndex[rightPointer].value
		if leftPointer == rightPointer {
			break
		}
		if sum == target {
			out[0] = numsWithIndex[leftPointer].originalindex
			out[1] = numsWithIndex[rightPointer].originalindex
			break
		}
		if sum < target {
			leftPointer++
		}
		rightPointer--

	}

	return out
}
