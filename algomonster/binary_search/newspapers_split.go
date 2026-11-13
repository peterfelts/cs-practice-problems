package binary_search

// Newspapers
// You have a stack of newspapers in a fixed order. Each newspaper has a read time.
// You want to assign all newspapers to a group of at most num_coworkers workers.
// Each worker is assigned a consecutive section of newspapers from the stack, and
// all workers read their assigned sections in parallel.

// The constraint: you cannot reorder newspapers. If you assign newspapers at
// positions 1, 2, 3 to worker A, you cannot then assign newspaper 2 to worker B.
// Each worker gets a consecutive block from the original stack.

// Find the minimum time needed to read all newspapers. Since workers read in
// parallel, the total time equals the time taken by the slowest worker.

// For example, with newspapers [7,2,5,10,8] and 2 workers, you could assign [7,2,5]
// to worker A (14 minutes total) and [10,8] to worker B (18 minutes total).
// Worker B finishes last, so the answer is 18 minutes.

// Constraints
// 1 <= newspapers_read_times.length <= 10^5
// 1 <= newspapers_read_times[i] <= 10^5
// 1 <= num_coworkers <= 10^5

// Examples
// Example 1:
// Input: newspapers_read_times = [7,2,5,10,8], num_coworkers = 2
// Output: 18
// Explanation:
// Assign first 3 newspapers to one coworker then assign the rest to another.
// The time it takes for the first 3 newspapers is 7 + 2 + 5 = 14 and for the
// last 2 is 10 + 8 = 18.

// Example 2:
// Input: newspapers_read_times = [2,3,5,7], num_coworkers = 3
// Output: 7
// Explanation:
// Assign [2, 3], [5], and [7] separately to workers. The minimum time is 7.

func newspapersSplit(newspapersReadTimes []int, numCoworkers int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	// The approach here involves simulating the greedy algorithm.
	// we will try to find the lowest time required to read all newspapers
	// by:
	//
	// 1) choose a target read-time value to test (say... 17 minutes)
	// 2) start at the beginning of the array, sum up read times
	// 3) If adding the N+1 item to the sum would exceed the target read
	//    time, add that item to the next reader.
	// 4) If you exceed the provided number of readers, that tells us that
	//    our target read time is too low, then we increase it. If we can
	//    read all the newspapers in that time, we try again and decrease
	//    the target read time (binary search).

	// First, let's find the bounds (the largest read time)
	maxInt := int(^uint(0) >> 1)
	shorestReadTime, largestReadTime, sumOfAllReadTimes := maxInt, 0, 0
	for _, readTime := range newspapersReadTimes {
		if readTime < shorestReadTime {
			shorestReadTime = readTime
		}

		if readTime > largestReadTime {
			largestReadTime = readTime
		}

		sumOfAllReadTimes += readTime
	}

	candidate := shorestReadTime

	left, right := largestReadTime, sumOfAllReadTimes

	// binary search
	for left <= right {
		mid := (left + right) / 2
		numReadersNeeded := 1

		// starting at the mid point, use the greedy algorithm
		// to count how many 'readers' are needed. If we exceed
		// the count provided, that tells us that mid is too
		// low, so we'll move right. If we can read all the
		// newspapers given the amount of readers, we move
		// left (lower mid point)
		sum := 0
		for _, readTime := range newspapersReadTimes {
			if sum+readTime <= mid {
				sum += readTime
			} else {
				numReadersNeeded++
				sum = readTime
			}
		}

		// We calculated that we need more readers than
		// we have provided coworkers, so we need to increase our
		// target read time
		if numReadersNeeded > numCoworkers {
			left = mid + 1
		} else {
			candidate = mid // possible shortest read time
			right = mid - 1
		}
	}

	return candidate
}
