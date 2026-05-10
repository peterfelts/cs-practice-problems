package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/1999209034/

// Given an integer array nums sorted in non-decreasing order, remove the duplicates
// in-place such that each unique element appears only once. The relative order of
// the elements should be kept the same.

// Consider the number of unique elements in nums to be k​​​​​​​​​​​​​​. After removing duplicates,
// return the number of unique elements k.

// The first k elements of nums should contain the unique numbers in sorted order.
// The remaining elements beyond index k - 1 can be ignored.

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	// use two pointers, l & r, to track source and destination indices
	l, r := 1, 1

	for {
		if r == len(nums) {
			return l
		}

		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
		r++
	}
}
