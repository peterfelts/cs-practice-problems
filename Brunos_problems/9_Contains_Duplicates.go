package bruno

// 9. Contains Duplicate (Arrays & Hash Sets)
// Problem: Given an integer array nums, return true if any value appears at least twice in the
// array, and return false if every element is distinct.

// Example:

// Input: nums = [1,2,3,1]

// Output: true

func containsDuplicates(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}
