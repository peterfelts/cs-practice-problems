package two_pointers

// Two Sum Sorted
// Given an array of integers sorted in ascending order, find two numbers that add
// up to a given target. Return the indices of the two numbers in ascending order.
// You can assume elements in the array are unique and there is only one solution.
// Do this in O(n) time and with constant auxiliary space.

// Input:

// arr: a sorted integer array
// target: the target sum we want to reach
// Sample Input: [2, 3, 4, 5, 8, 11, 18], 8

// Sample Output: 1 3

func twoSumSorted(arr []int, target int) []int {
	// WRITE YOUR BRILLIANT CODE HERE

	// two pointes:
	// left and right
	// on each iteration:
	// if L+R > target, move R to the left
	// if L+R < target, move L to the right

	left, right := 0, len(arr)-1

	for {

		sum := arr[left] + arr[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left, right}
		}
	}
}
