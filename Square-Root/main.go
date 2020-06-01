package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution(10, 20))
}
func Solution(a, b int) int {
	rs := 0
	a = int(math.Ceil(math.Sqrt(float64(a))))
	b = int(math.Floor(math.Sqrt(float64(b))))
	for b >= a {
		a = int(math.Ceil(math.Sqrt(float64(a))))
		b = int(math.Floor(math.Sqrt(float64(b))))
		rs++
	}
	return rs
}
