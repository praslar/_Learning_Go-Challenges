package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var results map[string]int

func main() {
	fmt.Println(Solution(1, 2, 3, 4))
}

func Solution(A int, B int, C int, D int) int {
	inputs := []int{A, B, C, D}
	temp := make([]int, 4)
	check := make([]bool, 4)
	results = make(map[string]int)
	BackTracking(inputs, temp, check, 0)
	fmt.Println(results)
	return len(results)
}

func BackTracking(inputs []int, temp []int, check []bool, index int) {
	for i, input := range inputs {
		if check[i] == false {
			check[i] = true
			temp[index] = input
			if index == 3 {
				time, ok := isTimeFormat(temp)
				if ok {
					results[time]++
				}
			} else {
				BackTracking(inputs, temp, check, index+1)
			}
			check[i] = false
		}
	}
}

func isTimeFormat(s []int) (string, bool) {
	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3])([0-9]|[0-5][0-9])$`)
	str := strconv.Itoa(s[0]) + strconv.Itoa(s[1]) + strconv.Itoa(s[2]) + strconv.Itoa(s[3])
	return str, re.MatchString(str)
}
