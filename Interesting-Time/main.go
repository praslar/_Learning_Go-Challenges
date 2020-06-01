package main

import (
	"fmt"
	"time"
)

func main() {
	s := "00:15:00"
	t := "00:16:00"
	fmt.Println(Solution(s, t))
}

func Solution(s string, t string) int {
	begin, _ := time.Parse("15:04:05", "00:00:00")
	start, _ := time.Parse("15:04:05", s)
	end, _ := time.Parse("15:04:05", t)
	rs := 0
	for i := start.Sub(begin); i <= end.Sub(begin); i += time.Second {
		if checkInter(i) {
			rs++
		}
	}
	return rs
}

func checkInter(t time.Duration) bool {
	check := map[string]int{}

	tmp := toString(t)
	for _, char := range tmp {
		check[string(char)]++
	}
	if len(check) < 3 {
		fmt.Println(tmp)
		fmt.Println(check)
	}
	return len(check) < 3

}

func toString(t time.Duration) string {
	tmp := int(t.Seconds())
	h := tmp / 3600
	m := (tmp % 3600) / 60
	s := (tmp % 3600) % 60
	return fmt.Sprintf("%02d%02d%02d", h, m, s)
}
