package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearchForStolenBike(t *testing.T) {

	fmt.Println()

	timeStolen := 23
	actual := is_stolen_bst(10, 100, timeStolen)

	require.Equal(t, timeStolen, actual)

	timeStolen = 12
	actual = is_stolen_bst(10, 100, timeStolen)
	require.Equal(t, timeStolen, actual)

}
