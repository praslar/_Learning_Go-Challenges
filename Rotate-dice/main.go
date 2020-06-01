package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{1, 6, 2, 3}))
}

const (
	MAX = 999
)

func Solution(dices []int) int {
	minRotate := MAX
	// IDEA:
	// Turn all dice to 1 -> get rotates move
	// Turn all dice to 2 -> get rotates move
	// ....
	// Turn all dive to 6 -> get rotates move
	// return the smallest rotates move

	// Number of pips on each each face of a dice
	for pips := 1; pips <= 6; pips++ {

		rotate := 0
		// With each pips, rotate all dices to this pip
		for side := 0; side < len(dices); side++ {
			// Rotate 2
			if pips+dices[side] == 7 {
				rotate += 2
				// Rotate 1
			} else if pips != dices[side] {
				rotate++
			}
		}
		// get the smallest rotate time
		if rotate < minRotate {
			minRotate = rotate
		}

	}

	return minRotate
}
