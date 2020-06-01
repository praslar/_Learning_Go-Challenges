package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "We test coders. Give us a try?"

	fmt.Println(Solution(s))
}

func Solution(s string) int {
	ns := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.' || r == '!' || r == '?'
	})
	max := 0
	for _, v := range ns {
		ws := strings.FieldsFunc(v, func(r rune) bool {
			return unicode.IsSpace(r)
		})
		if max < len(ws) {
			max = len(ws)
		}
	}
	return max
}
