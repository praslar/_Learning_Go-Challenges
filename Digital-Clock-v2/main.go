package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(Solution([]int{1, 8, 3, 2}))
}

var results map[string]int

func Solution(inputs []int) int {

	check := make([]bool, len(inputs))
	tmp := make([]int, 4)
	results = map[string]int{}
	BackTracking(check, inputs, 0, tmp)
	return len(results)
}

func BackTracking(check []bool, inputs []int, count int, tmp []int) {
	for index, input := range inputs {
		if check[index] == false {
			check[index] = true
			tmp[count] = input
			if count == 3 {
				if time := isTimeFormat(tmp); time != "" {
					results[time]++
				}
			} else {
				BackTracking(check, inputs, count+1, tmp)
			}
			check[index] = false
		}
	}
}

func isTimeFormat(tmp []int) string {
	input := strconv.Itoa(tmp[0]) + strconv.Itoa(tmp[1]) + strconv.Itoa(tmp[2]) + strconv.Itoa(tmp[3])
	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3])([0-9]|[0-5][0-9])$`)
	if re.MatchString(input) {
		fmt.Println("valid time", input)
		return input
	}
	return ""
}
