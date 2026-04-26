package main

import (
	"testing"
)

var jetpackTests = []struct {
	initialVelocity int
	surface         []int
	expected        bool
}{
	{
		5,
		[]int{1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0},
		true,
	},
	// initialVelocity=7, strictly decelerating by 1 each hop:
	// 0 --(+6)--> 6 --(+5)--> 11 --(+4)--> 15 --(+3)--> 18 --(+2)--> 20 --(+1)--> 21 --(+0)--> 21(stop)
	// safe positions: 0, 6, 11, 15, 18, 20, 21
	{
		7,
		[]int{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1},
		true,
	},
	// initialVelocity=4, safe spots ONLY at 0, 3, 5, 9, 14 (all others are craters).
	// Every reachable path either lands on a crater or exits the array before speed=0:
	//   (0,4)->(3,3)->(5,2)->{6,7,8} all craters
	//   (0,4)->(4,4)           crater
	//   (0,4)->(5,5)->(9,4)->{12,13} craters
	//   (0,4)->(5,5)->(9,4)->(14,5)->{18,19,20} craters / OOB
	//   (0,4)->(5,5)->{10,11}   craters
	{
		4,
		[]int{1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		false,
	},
}

func Test_Jump(t *testing.T) {

	// note: speed can change +/- 1, or stay constant

	for _, test := range jetpackTests {
		actual := jump(test.initialVelocity, 0, test.surface)

		if actual != test.expected {
			t.Fail()
		}
	}

}

func Test_JumpWithMemorization(t *testing.T) {
	for _, test := range jetpackTests {
		// Reset global memoization state between test cases to prevent
		// failed states from one surface poisoning the next.
		jumpMemory = make(map[step]struct{})
		actual := jumpWithMemorization(test.initialVelocity, 0, test.surface)

		if actual != test.expected {
			t.Fail()
		}
	}
}
