// To execute Go code, please declare a func main() in a package "main"

package main

// Lunar Jetpack
// An engineer is traversing the moon's surface from one lunar office to another using a jetpack to control their
// movement. The path between the offices is scattered with hazardous craters that the engineer must avoid.
// Rules:
// Lunar Path: The path is represented by a boolean array, where each element corresponds to a discrete spot on the
// moon's surface. A value of True indicates a safe spot, while False represents a hazardous crater.
// Starting Speed (S): The engineer starts with a non-negative integer speed S. This speed determines how many spots
// forward they will move on their next jetpack burst.
// Adjusting Speed: After each landing (including the initial position), the engineer can adjust their speed by
// increasing or decreasing it by 1 unit (or keeping it the same) before the next move. This adjustment simulates
// controlling the jetpack's thrust.
// Goal: The engineer must safely come to a complete stop (speed becomes 0) before reaching the other lunar office
// (the end of the path). Failing to stop in time means crashing into the office dome, resulting in mission failure.
// Hazards: Landing on a hazardous crater at any point causes the engineer's jetpack to malfunction, and the journey
// ends in failure.

// Objective:
// Develop a function that returns a boolean value indicating whether the engineer can safely come to a stop somewhere
// along the lunar path—without ever landing on a hazardous spot and before reaching the other lunar office.

func jump(speed, i int, path []int) bool {

	// terminal conditions:
	// end could be valid point
	if i >= len(path) {
		return false
	}

	if path[i] == 0 {
		return false
	}

	if speed == 0 {
		return true
	}

	// TODO: we're at the end and it's true || false

	indexDecrease := i + speed - 1
	indexIncrease := i + speed + 1

	// decrease || constant || increase
	return jump(speed-1, indexDecrease, path) || jump(speed, i+speed, path) || jump(speed+1, indexIncrease, path)

}

// Data type used to store details of a step
// that has been taken
type step = struct {
	index int
	speed int
}

var jumpMemory = make(map[step]struct{})

func jumpWithMemorization(speed, i int, path []int) bool {

	// Check if we have already been to this index with an identical
	// speed
	if _, ok := jumpMemory[step{i, speed}]; ok {
		return false
	}

	// store this jump in memory
	jumpMemory[step{i, speed}] = struct{}{}

	// terminal conditions:
	// end could be valid point
	if i >= len(path) {
		return false
	}

	if path[i] == 0 {
		return false
	}

	if speed == 0 {
		return true
	}

	// TODO: we're at the end and it's true || false

	indexDecrease := i + speed - 1
	indexIncrease := i + speed + 1

	// decrease || constant || increase
	return jump(speed-1, indexDecrease, path) || jump(speed, i+speed, path) || jump(speed+1, indexIncrease, path)

}

// https://docs.google.com/document/d/1a4RmEPJV9G0bLMr16U0KCJXDp8h-ttkfcf9PTawdzR8/edit?tab=t.0
