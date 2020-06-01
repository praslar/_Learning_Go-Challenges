package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{100, 200, 100}, []int{50, 100, 100}, 100, 100))
}

func Solution(A []int, B []int, X int, Y int) int {
	tapArea := 20 * 20
	for i := 0; i < len(A); i++ {
		if distance(A[i], B[i], X, Y) <= tapArea {
			return i
		}
	}
	return -1
}

func distance(xA int, yA int, xB int, yB int) int {
	return ((xA-xB)*(xA-xB) + (yA-yB)*(yA-yB))
}
