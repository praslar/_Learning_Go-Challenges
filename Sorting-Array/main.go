package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{1, 5, 3, 3, 7}))
}
func Solution(a []int) bool {

	if sort.IntsAreSorted(a) {
		return true
	}
	// Calculate how many time need to swap numbers
	countDif := 0

	sortedArr := append([]int{}, a...)

	sort.Ints(sortedArr)
	// For each pair of number between to arrs
	for i := 0; i < len(a); i++ {
		if a[i] != sortedArr[i] {
			countDif++
			if countDif > 2 {
				return false
			}
		}
	}
	// If need to swap more than 2 number
	// Return false
	return true
}
