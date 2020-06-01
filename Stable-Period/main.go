package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{1, 3, 5, 9, 10, 92}))
}

func Solution(A []int) int {
	//TODO your code
	rs := 0
	for i := 0; i < len(A)-1; i++ {
		count := 0
		stable := 0
		for j := i + 1; j < len(A); j++ {
			count++
			temp := A[j] - A[j-1]
			if count >= 2 {
				if temp == stable {
					rs++
				} else {
					break
				}
			}
			stable = temp
		}
	}
	return rs
}
