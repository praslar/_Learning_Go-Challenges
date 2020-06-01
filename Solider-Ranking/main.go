package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{3, 4, 3, 0, 2, 2, 3, 0, 0}))
}

func Solution(ranks []int) int {
	result := 0
	solider := map[int]int{}
	for i := 0; i < len(ranks); i++ {
		// điểm danh nếu cùng cấp bậc
		solider[ranks[i]]++
	}
	for k, v := range solider {
		_, found := solider[k+1]
		if found {
			result += v
		}
	}
	return result
}
