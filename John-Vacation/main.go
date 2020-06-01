package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(Solution(2020, "March", "March", "April"))
}
func Solution(year int, start string, end string, w string) int {
	WeekDays := 7
	TimeLayout := "2006-January-2"
	MonthsIndex := map[string]int{"January": 0, "February": 1, "March": 2, "April": 3, "May": 4, "June": 5, "July": 6, "August": 7, "September": 8, "October": 9, "November": 10, "December": 11}
	DayOfMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%4 == 0 {
		DayOfMonths[1] = 29
	}

	startDay, _ := time.Parse(TimeLayout, format(year, start, 1))
	endDay, _ := time.Parse(TimeLayout, format(year, end, DayOfMonths[MonthsIndex[end]]))

	goaway := int(startDay.Weekday())
	comehome := int(endDay.Weekday())
	total := 0

	for i := MonthsIndex[start]; i <= MonthsIndex[end]; i++ {
		fmt.Println(i)
		total += DayOfMonths[i]
	}

	if goaway == 0 {
		total--
	}
	if goaway != 1 && goaway != 0 {
		total -= (WeekDays - goaway + 1)
	}
	fmt.Println(total)
	if comehome != 0 {
		total -= comehome
	}
	fmt.Println(total)

	return total / 7
}

func format(y int, m string, d int) string {
	return fmt.Sprintf("%v-%v-%v", y, m, d)
}
