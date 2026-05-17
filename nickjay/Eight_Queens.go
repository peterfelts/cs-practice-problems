package nickjay

import "fmt"

// queen can move diagnoal, horizontal, vertical
// board 8x8 -> 64
// Given 8 queens, find all layouts that can fit all 8 queens,
// without any of them being in a position where they can capture
// another.

// output: array of data structures used to represent positions

// Given a position of a queen:
// - no other queen in the same row or column
// - no other queen in any diagonal

// 0

// TODO: tuple isn't needed because the index in the solution array can represent the row
// print out all solutions

const EIGHT_QUEENS int = 8

func eightQueens() [][]int {

	positions := []int{} // the index will be the row, the value will be the column
	allSolutions := [][]int{}

	allSolutions = populate(0, positions, allSolutions)

	for i, solution := range allSolutions {
		fmt.Printf("Found %d solutions. Solution #%d\n", len(allSolutions), i+1)
		printBoard(solution)
		fmt.Println()
	}

	return allSolutions
}

func populate(row int, positions []int, allSolutions [][]int) [][]int {

	// we've reached the end of the board
	if row == EIGHT_QUEENS {

		// if we've reached past the final row, we have a solution
		// copy the solution into a new slice, otherwise, any changes
		// we make further back in the stack will influence this memory
		tmp := make([]int, len(positions))
		copy(tmp, positions)
		allSolutions = append(allSolutions, tmp)

		// NOTE: be sure not to return 'tmp' here, as that would overwrite
		// the caller's slice
		return allSolutions
	}

	for j := 0; j < EIGHT_QUEENS; j++ { // iterate over each column

		// if we can place a queen in this position,
		// recursiverly call into the next row
		if canPlace(row, j, positions) {

			// place the queen in the current position temporarily.
			// Remove it so that we can test other permutations
			positions = append(positions, j)

			allSolutions = populate(row+1, positions, allSolutions)

			// remove the current queen
			positions = positions[:row]

		}
	}

	return allSolutions
}

// Given the current row and column, determine if there are any queens
// present in the same row, column, or diagonals
func canPlace(row, col int, positions []int) bool {

	// cycle through all previous positions
	for i := 0; i < len(positions); i++ {
		// check row and column
		// i == row -> there is already a queen in the current row (sholdnt' happen)
		// positions[i] == col -> there is a queen in the same column
		if i == row || positions[i] == col {
			return false
		}
	}

	// check diagonals
	// left diagonal
	i := row - 1
	j := col - 1
	for {
		if i < 0 || j < 0 {
			break
		}

		// we've found a queen on the diagonal
		if positions[i] == j {
			return false
		}
		i--
		j--
	}

	// right diagonal
	i = row - 1
	j = col + 1
	for {
		if i < 0 || j > EIGHT_QUEENS-1 {
			break
		}

		if positions[i] == j {
			return false
		}

		i--
		j++

	}

	return true
}

// printBoard prints an ASCII representation of an 8x8 board given a positions slice,
// where the index is the row and the value is the column of the queen in that row.
// Q marks a queen, . marks an empty square.
func printBoard(positions []int) {
	fmt.Println("  0 1 2 3 4 5 6 7")
	for row := 0; row < EIGHT_QUEENS; row++ {
		fmt.Printf("%d ", row)
		for col := 0; col < EIGHT_QUEENS; col++ {
			if row < len(positions) && positions[row] == col {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
