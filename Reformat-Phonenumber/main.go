package main

import "fmt"

func main() {
	fmt.Println(Solution("00-44 48 5555 8361"))
}
func Solution(s string) string {
	number := make([]rune, 0)
	//Get only number from string
	for _, char := range s {
		if char != '-' && char != ' ' {
			number = append(number, char)
		}
	}
	// Add dashes
	s = string(number)
	// Return "000"
	if len(s) < 3 {
		return s
	}
	// Return "12-21"
	if len(s) == 4 {
		return s[:len(s)-2] + "-" + s[len(s)-2:]
	}
	// Return "333-999-111"
	if len(s)%3 == 0 {
		for i := 3; i < len(s); i += 4 {
			s = s[:i] + "-" + s[i:]
		}
		return string(s)
	}
	if len(s)%3 == 2 {
		//Return "333-333-12"
		for i := 3; i < len(s)-2; i += 4 {
			s = s[:i] + "-" + s[i:]
		}
		s = s[:len(s)-2] + "-" + s[len(s)-2:]
		return string(s)
	}
	if len(s)%3 == 1 {
		// Return "333-111-12-31"
		for i := 3; i < len(s)-4; i += 4 {
			s = s[:i] + "-" + s[i:]
		}
		s = s[:len(s)-4] + "-" + s[len(s)-4:]
		s = s[:len(s)-2] + "-" + s[len(s)-2:]
		return string(s)
	}
	return string(s)
}
