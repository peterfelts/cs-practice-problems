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

func TestEightQueens(t *testing.T) {
	eightQueens()
}
