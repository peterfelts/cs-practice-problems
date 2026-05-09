package main

// Given an array of non-negative integers height representing an elevation map,
// compute how much water it can trap after raining.
//
// For example:
// The array: [0,1,0,5,1,5] can be visualized as:
//
//             +---+   +---+
//             |   |   |   |
//             |   |   |   |
//             |   |   |   |
//     +---+   |   +---+   |
//     |   |   |   |   |   |
// ----+---+---+---+---+---+----
//  0   1   0   5   1   5
//
// For this case, the max volume of water that could be captured
// between these rectangles is 1 + 3 = 4

func maxWaterVolumeBruteForce(heights []int) int {

	maxWaterVolume := 0

	// our algorithm will be:
	// at each rectangle, the volume (area) of water that each
	// rectangle can hold is min(maxL, maxR) - heigt
	for i := 0; i < len(heights)-1; i++ {

		maxL, maxR := 0, 0 // the maxium height to the left and reight of each index

		// O(n^2) brute-force approach: recalculate maxL, maxR
		// on each iteration
		for k := 0; k < len(heights); k++ {

			// look left
			if k < i && heights[k] > maxL {
				maxL = heights[k]
			}
			// look right
			if k > i && heights[k] > maxR {
				maxR = heights[k]
			}
		}

		volume := min(maxL, maxR) - heights[i]
		if volume > 0 {
			maxWaterVolume += min(maxL, maxR) - heights[i]
		}
	}

	return maxWaterVolume
}

func maxWaterVolumeDP(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	maxWaterVolume := 0
	maxL, maxR := heights[0], heights[0] // the maxium height to the left and reight of each index

	for i := 0; i < len(heights)-1; i++ {

		// look left
		if heights[i] > maxL {
			maxL = heights[i]
		}
		// Find a new maxR once we've
		// progressed to and reached the current maxR
		if heights[i] == maxR {
			maxR = 0
			for k := i + 1; k < len(heights); k++ {
				if heights[k] > maxR {
					maxR = heights[k]
				}
			}
		}

		volume := min(maxL, maxR) - heights[i]
		if volume > 0 {
			maxWaterVolume += min(maxL, maxR) - heights[i]
		}
	}

	return maxWaterVolume
}
