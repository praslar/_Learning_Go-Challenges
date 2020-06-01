package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	t := []string{"test1a", "test2", "test1b"}
	r := []string{"Wrong answer", "OK", "Runtime error"}
	fmt.Println(Solution(t, r))
}

func Solution(t []string, r []string) int {
	//init value
	mapRs := make(map[string][]string)

	for i := 0; i < len(t); i++ {
		testGroup := split(t[i])
		mapRs[testGroup] = append(mapRs[testGroup], r[i])
	}

	point := len(mapRs)
	for i, group := range mapRs {
		for _, result := range group {
			if result != "OK" {
				mapRs[i] = []string{"Fail"}
				point--
				break
			}
		}
	}

	return point * 100 / len(mapRs)
}

func split(s string) string {
	f := func(c rune) bool {
		return !unicode.IsDigit(c)
	}
	return strings.FieldsFunc(s, f)[0]
}
