package main

import "fmt"

func main() {
	trees := []int{6, 1, 4, 6, 3, 2, 7, 4}
	fmt.Println(Solution(trees, 3, 2))
}

func Solution(trees []int, Alice int, Bob int) int {

	if Alice+Bob > len(trees) {
		return -1
	}

	aliceMax, aliceSegment := FindSegment(trees, Alice, 0, 0)

	remain := RemoveTrees(trees, aliceSegment, Alice)

	bobMax, _ := FindSegment(remain, Bob, 0, 0)

	return aliceMax + bobMax
}

func RemoveTrees(a []int, i int, alice int) []int {
	copy(a[i:], a[i+alice:])
	a = a[:len(a)-alice]
	return a
}

func FindSegment(trees []int, quantity int, apples int, segment int) (int, int) {
	for i := 0; i <= len(trees)-quantity; i++ {
		max := 0
		for j := 0; j < quantity; j++ {
			max += trees[i+j]
		}
		if max > apples {
			apples = max
			segment = i
		}
	}
	return apples, segment
}
