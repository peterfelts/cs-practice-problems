package two_pointers

// Container With Most Water

// You are given an array height where height[i] is the height of a vertical bar at
// position i. All bars are spaced one unit apart. Two bars and the x-axis form a
// container that can hold water. The amount of water is
// (j - i) × min(height[i], height[j]): the width between the bars multiplied by
// the shorter bar's height, since water spills over the shorter side.

// Find the pair of bars that holds the most water.

// Input: [1, 8, 6, 2, 5, 4, 8, 3, 7]

// Output: 49

// The best pair is bars at index 1 and index 8, with heights 8 and 7.
// Width is 8 - 1 = 7, shorter height is 7, so the area is 7 × 7 = 49.

func containerWithMostWater(arr []int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	// two pointers: L and R
	// On each iteration, calculate the volume: (R-L) x min(height[L], height[R])
	// If height[L] < height[R], the only way to possibly get a larger volume is to move L to the right
	// if height[L] > height[R], the only way to possibly get a larger volume is to move R to the left
	left, right := 0, len(arr)-1
	maxVolume := 0

	for left < right {
		volume := (right - left) * min(arr[left], arr[right])
		if volume > maxVolume {
			maxVolume = volume
		}

		// keep going
		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolume
}
