package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(1, 111))
}

func Solution(from, to int) int {
	rs := 0
	for i := from; i < to; i++ {
		if isLovely(i) {
			rs++
		}
	}
	return rs
}
func isLovely(numb int) bool {
	digits := strconv.Itoa(numb)
	checking := map[rune]int{}
	for _, digit := range digits {
		checking[digit]++
		if checking[digit] > 2 {
			return false
		}
	}
	return true
}
