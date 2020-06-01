package main

import "fmt"

func main() {

	bulbs := []int{5, 3, 1, 2, 4}
	fmt.Println(Solution(bulbs))
}

func Solution(bulbs []int) int {
	rs := 0
	light := make([]int, 0)
	max := 0
	for i := 0; i < len(bulbs); i++ {
		light = append(light, bulbs[i])
		if bulbs[i] > max {
			max = bulbs[i]
		}
		if len(light) == max {
			rs++
		}
		fmt.Println(light)
	}
	return rs
}
