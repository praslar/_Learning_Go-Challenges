package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	b := `Sun 10:00-20:00
Fri 05:00-10:00 
Fri 16:30-23:50 
Sat 10:00-24:00 
Sun 01:00-04:00 
Sat 02:00-06:00 
Tue 03:30-18:15 
Tue 19:00-20:00
Wed 04:25-15:14
Wed 15:14-22:40
Thu 00:00-23:59
Mon 05:00-13:00
Mon 15:00-21:00`
	fmt.Println(Solution(b))
}

const (
	secondPerDay = 24 * 60
)

func Solution(b string) int {
	meeting := getWorkTime(b)
	MAX := 0
	for i := len(meeting); i > 0; i-- {
		if i%2 != 0 && meeting[i]-meeting[i-1] > MAX {
			MAX = meeting[i] - meeting[i-1]
		}
	}
	return MAX
}

func getWorkTime(s string) []int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != ':'
	}
	split := strings.FieldsFunc(s, f)

	workTime := map[string][]int{}

	for i := 0; i < len(split)-2; i += 3 {
		DayOfWeek := split[i]
		start := split[i+1]
		end := split[i+2]
		workTime[split[i]] = append(workTime[DayOfWeek], timeToInt(start, DayOfWeek), timeToInt(end, DayOfWeek))
	}

	meeting := []int{0, secondPerDay * 7}
	for _, time := range workTime {
		meeting = append(meeting, time...)
	}
	sort.Ints(meeting)
	return meeting
}

func timeToInt(t, d string) int {

	dayofweek := map[string]int{
		"Mon": 0,
		"Tue": secondPerDay,
		"Wed": secondPerDay * 2,
		"Thu": secondPerDay * 3,
		"Fri": secondPerDay * 4,
		"Sat": secondPerDay * 5,
		"Sun": secondPerDay * 6,
	}

	a, _ := time.Parse("04:05", t)
	b, _ := time.Parse("04:05", "00:00")

	return dayofweek[d] + int(a.Sub(b).Seconds())
}
