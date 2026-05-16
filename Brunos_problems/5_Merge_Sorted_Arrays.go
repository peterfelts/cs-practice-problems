package bruno

import "fmt"

// 5. Merge Sorted Array (Arrays & Two Pointers)
// Problem: You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of valid elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order. The final sorted array
// should be stored inside the array nums1. The length of nums1 is m + n.

// Example:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3

// Output (nums1 should be): [1,2,2,3,5,6]

func mergeSortedArrays(nums1, nums2 []int, m, n int) []int {

	// use two pointers, each pointing at the maximum value
	// in each array
	nums1Index := m - 1
	nums2Index := n - 1

	fmt.Println(nums1)
	fmt.Println(nums2)

	// We know that each array is already sorted and we know
	// that we have a lot of empty space after the m-1 index
	// in the first array. We'll walk the length of that array,
	// from right to left and we'll whichever value is greater
	// at index i
	for i := m + n - 1; i >= 0; i-- {

		// nums1Index < 0 -> copy from nums2
		// nums2Index < 0 -> copy from nums1
		// nums1Index >= 0 && nums2Index >= 0 -> figure out which to copy
		if nums1Index < 0 {
			nums1[i] = nums2[nums2Index]
			nums2Index--
		} else if nums2Index < 0 {
			nums1[i] = nums1[nums1Index]
			nums1Index--
		} else if nums2[nums2Index] > nums1[nums1Index] {
			nums1[i] = nums2[nums2Index]
			nums2Index--
		} else {
			nums1[i] = nums1[nums1Index]
			nums1Index--
		}
	}

	return nums1
}
