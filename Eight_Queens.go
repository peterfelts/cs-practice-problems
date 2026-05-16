package main

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

	_, positions = populate(0, positions)

	printBoard(positions)

	return [][]int{}
}

func populate(row int, positions []int) (bool, []int) {

	// we've reached the end of the board
	if row == EIGHT_QUEENS {
		return true, positions
	}

	branchWorks := false
	for j := 0; j < EIGHT_QUEENS; j++ { // iterate over each column

		branchWorks = false

		// if we can place a queen in this position,
		// recursiverly call into the next row
		if canPlace(row, j, positions) {

			// place the queen in the current position temporarily.
			// Remove it if we can't find a suitable path
			positions = append(positions, j)

			branchWorks, positions = populate(row+1, positions)
			if branchWorks {

				// append to the positions
				positions = append(positions, j)
				break
			}
			// pop
			positions = positions[:row]
		}

	}

	return branchWorks, positions
}

// [row, col] = [0,0]

// Given the current row and column, determine if there are any queens
// present in the same row, column, or diagonals
func canPlace(row, col int, positions []int) bool {

	// cycle through all previous positions
	for i := 0; i < len(positions); i++ {
		// check row and column
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
		fmt.Printf("checking diagonal: %d, %d, row, col: (%d,%d)\n", i, positions, row, col)
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
		if i < 0 || col > EIGHT_QUEENS-1 {
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
