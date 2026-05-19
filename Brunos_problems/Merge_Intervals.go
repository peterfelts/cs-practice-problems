package bruno

import (
	"fmt"
	"sort"
)

// Given an array of intervals where intervals[i] = [start, end], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.

// [[1, 3] [2, 5]] -> [[1, 5]]
// [[1, 3], [2, 5], [2, 100]] -> [1, 100]
// [[1, 6], [4, 5], [2, 100]]

// [[1, 100]]

// assume:
//  int32
//  start <= end [3, 3]
//  no malformed intervals
//  intervals are not sorted, arbitrary
//  no negative values
//  overlapping: any intersection of two intervals

type interval struct {
	Left  int
	Right int
}

// approach:
// 1) sort the array by left value
// 2) iterate through the array, comparing intervals to detect overlap.
//   - compare the most recent interval placed in the merged output slice
//     with the current interval.
//   - if an overlap is detected, update the last interval in the merged
//     slice.
//   - if no overlap is detected, add the interval to the end of the
//     merged slice
func mergeIntervals(intervals []interval) []interval {

	if len(intervals) < 2 {
		return intervals
	}

	fmt.Println(intervals)

	// Sort intervals
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Left < intervals[b].Left
	})

	fmt.Println(intervals)

	merged := []interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		l := merged[len(merged)-1]
		r := intervals[i]

		// detect overlap
		if l.Right >= r.Left {
			merged[len(merged)-1] = interval{
				min(l.Left, r.Left),
				max(l.Right, r.Right),
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	fmt.Println(merged)

	return merged
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
