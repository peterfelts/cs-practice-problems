// To execute Go code, please declare a func main() in a package "main"

package main

// # Max-Sum Path

// Given a non-empty grid of positive integers, `grid`, find the path from the top-left corner to the bottom-right corner with the largest sum. You can only go down or to the right (not diagonal).

// Example 1: grid = [[1, 4, 3],
//                    [2, 7, 6],
//                    [5, 8, 9]]
// Output: 29. The maximum path is 1 -> 4 -> 7 -> 8 -> 9.

// Example 2: grid = [[5]]
// Output: 5

// Example 3: grid = [[1, 2, 3]]
// Output: 6. The maximum path is 1 -> 2 -> 3.

// Constraints:

// - `1 <= R, C <= 1000`, where `R` is the number of rows and `C` is the number of columns in the grid.
// - `1 <= grid[i][j] <= 1000`.

func findMaxSumPath(grid [][]int, rIndex, cIndex, r, c int, memo map[[2]int]int) int {

	// return this branch if it's already been calculated
	if branchValue, ok := memo[[2]int{rIndex, cIndex}]; ok {
		return branchValue
	}

	// if we can't go any further
	if rIndex == r-1 && cIndex == c-1 {
		return grid[r-1][c-1]
	}

	// If we're not at the bottom of the array
	down := 0
	if rIndex < r-1 {
		down = findMaxSumPath(grid, rIndex+1, cIndex, r, c, memo)
	}

	right := 0
	if cIndex < c-1 {
		right = findMaxSumPath(grid, rIndex, cIndex+1, r, c, memo)
	}

	// store the calculated branch value
	branchValue := grid[rIndex][cIndex] + max(down, right)
	memo[[2]int{rIndex, cIndex}] = branchValue

	return branchValue

}
