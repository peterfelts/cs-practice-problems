package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlace(t *testing.T) {
	board := []int{0} // a queen at position 0,0

	// left diagonal - expected to fail
	actual := canPlace(1, 1, board)
	assert.False(t, actual)

	// test vertical - expected to fail
	board = []int{0, 2}
	assert.False(t, canPlace(0, 0, board))

	// right diagonal - expected to fail
	board = []int{3}
	assert.False(t, canPlace(1, 2, board))

	// expected to pass
	board = []int{0}
	assert.True(t, canPlace(1, 2, board))
}

// TestEightQueensSolutionCount verifies that exactly 92 solutions exist,
// which is the well-known answer for the 8 queens problem.
func TestEightQueensSolutionCount(t *testing.T) {
	solutions := eightQueens()
	assert.Equal(t, 92, len(solutions), "expected exactly 92 solutions")
}

// TestEightQueensSolutionLength verifies that every solution has exactly 8 queens.
func TestEightQueensSolutionLength(t *testing.T) {
	solutions := eightQueens()
	for i, s := range solutions {
		assert.Equal(t, 8, len(s), "solution %d has wrong length", i)
	}
}

// TestEightQueensNoColumnConflict verifies that no two queens in any solution
// share the same column.
func TestEightQueensNoColumnConflict(t *testing.T) {
	solutions := eightQueens()
	for i, s := range solutions {
		seen := map[int]bool{}
		for _, col := range s {
			assert.False(t, seen[col], "solution %d has duplicate column %d", i, col)
			seen[col] = true
		}
	}
}

// TestEightQueensNoDiagonalConflict verifies that no two queens in any solution
// share a diagonal.
func TestEightQueensNoDiagonalConflict(t *testing.T) {
	solutions := eightQueens()
	for i, s := range solutions {
		for r1 := 0; r1 < len(s); r1++ {
			for r2 := r1 + 1; r2 < len(s); r2++ {
				rowDiff := r2 - r1
				colDiff := s[r2] - s[r1]
				if colDiff < 0 {
					colDiff = -colDiff
				}
				assert.NotEqual(t, rowDiff, colDiff,
					"solution %d: queens at rows %d and %d share a diagonal", i, r1, r2)
			}
		}
	}
}

// TestEightQueensUniqueSolutions verifies that no two solutions are identical.
func TestEightQueensUniqueSolutions(t *testing.T) {
	solutions := eightQueens()
	seen := map[[8]int]bool{}
	for i, s := range solutions {
		var key [8]int
		copy(key[:], s)
		assert.False(t, seen[key], "solution %d is a duplicate", i)
		seen[key] = true
	}
}

// TestEightQueensColumnsInRange verifies that every column value is within 0–7.
func TestEightQueensColumnsInRange(t *testing.T) {
	solutions := eightQueens()
	for i, s := range solutions {
		for row, col := range s {
			assert.True(t, col >= 0 && col < 8,
				"solution %d, row %d: column %d is out of range", i, row, col)
		}
	}
}
