package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("Baaa"))
}

func Solution(S string) string {

	rs := []rune{}
	caplocks := false

	for _, char := range S {
		if char == 'C' {
			caplocks = !caplocks
			continue
		}
		if char == 'B' {
			if len(rs) > 0 {
				rs = rs[:len(rs)-1]
			}
			continue
		}
		if caplocks {
			rs = append(rs, toUpperCase(char))
			continue
		}
		rs = append(rs, char)
	}
	return string(rs)
}

func toUpperCase(s rune) rune {
	return s - 32
}
