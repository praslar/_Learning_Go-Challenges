package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	secondPerDay = 24 * 60
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

//Solution jafar
func Solution(b string) int {
	m := strings.Split(b, "\n")

	n := make(map[int][]string)

	for _, v := range m {
		parts := strings.Split(v, " ")
		each := strings.Split(parts[1], "-")

		switch parts[0] {
		case "Mon":
			n[0] = append(n[0], each[0], each[1])
		case "Tue":
			n[1] = append(n[1], each[0], each[1])
		case "Wed":
			n[2] = append(n[2], each[0], each[1])
		case "Thu":
			n[3] = append(n[3], each[0], each[1])
		case "Fri":
			n[4] = append(n[4], each[0], each[1])
		case "Sat":
			n[5] = append(n[5], each[0], each[1])
		case "Sun":
			n[6] = append(n[6], each[0], each[1])
		}

	}

	meeting := []int{0, secondPerDay * 7}

	for i, val := range n {
		for _, v := range val {
			each := strings.Split(v, ":")
			h, _ := strconv.Atoi(each[0])
			m, _ := strconv.Atoi(each[1])
			meeting = append(meeting, (i*1440)+(h*60+m))
		}
	}

	sort.Ints(meeting)
	for i, v := range meeting {
		fmt.Println(i, v)
	}

	sleep := []int{}
	for i := len(meeting); i > 0; i-- {
		if i%2 != 0 {
			sleep = append(sleep, meeting[i]-meeting[i-1])
		}
	}
	for i, v := range sleep {
		fmt.Println(i, v)
	}
	sort.Ints(sleep)
	return sleep[len(sleep)-1]
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
