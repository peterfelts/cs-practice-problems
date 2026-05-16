package leetcode

// https://leetcode.com/problems/contiguous-array/description/

// I followed this video to get to an O(n) time-complexity solution. I would not have
// figured this out on my own within the timeframe of an interview. (I did come up with
// an O(n^2) time-complexity solution, but it was not accepted because the tests took too long)
//
// https://www.youtube.com/watch?v=agB1LyObUNE
//
// The approach is to do a pass through the array and count the difference between the
// count of ones and the count of zeros at each point in the array:
//      Array: [ 0 1 1 1 1 1 0 0 0]
// Diff Array: [-1 0 1 2 3 4 3 2 1]
//
// The "Diff Array" will serve as a map between a given diff (count of 1's - count of 0's)
// and the first index in the array where this diff value was seen. We'll use a map for
// this and will not replace/update existing keys, because we always want the map to store
// the lowest index that contains the diff we're looking for at any point in time.
//
// Then, when we go through the input array and keep a running diff, at each point we look
// for an earlier index in the Diff Map with the same diff value. If found, the subarray
// between that index and the current index has equal ones and zeros.

func findMaxLength(nums []int) int {
	diffMap := make(map[int]int, len(nums))

	// First pass through the array. Populate the diff map
	diff := 0
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			diff++
		} else {
			diff--
		}

		// only insert key if it doesn't exist
		if _, ok := diffMap[diff]; !ok {
			diffMap[diff] = i
		}

		// special case (we got lucky)
		if diff == 0 {
			if i > maxLength {
				maxLength = i + 1
			}
		}

		// fmt.Printf("i: %d, diff: %d\n", i, diff)
		// At each point, see if there is an ideal diff such that the running diff - ideal diff == 0.
		// If there is, then we calculate the subarray length and store it if is greater than
		// any previously-calculated length
		if startIndex, ok := diffMap[diff]; ok {
			subArrayLength := i - startIndex
			if subArrayLength > maxLength {
				maxLength = subArrayLength
			}
		}
	}

	return maxLength
}
