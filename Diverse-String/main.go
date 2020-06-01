package main

import "fmt"

func main() {
	a, b, c := 8, 8, 2
	fmt.Println(Solution(a, b, c))
}
func Solution(a, b, c int) string {
	var result []byte
	sum := a + b + c
	for i := 0; i < sum; i++ {
		if a > 0 {
			result = insertToString('a', result)
			a--
		}
		if b > 0 {
			result = insertToString('b', result)
			b--
		}
		if c > 0 {
			result = insertToString('c', result)
			c--
		}
	}
	return string(result)
}
func insertToString(b byte, s []byte) []byte {
	if len(s) < 3 {
		return append(s, b)
	}
	for i := 0; i < len(s); i++ {
		// insert first
		fmt.Println(string(s))
		if i == 0 && s[i] != s[i+1] {
			s = append(s, 0)
			copy(s[i+1:], s[i:])
			s[i] = b
			return s
		}
		//insert last
		if i == len(s)-1 && b != s[i] || i == len(s)-1 && b != s[i-1] {
			return append(s, b)
		}
		// insert mid
		if i < len(s)-2 && b != s[i] && b != s[i+1] || i < len(s)-2 && b != s[i] && s[i+1] != s[i+2] {
			s = append(s, 0)
			copy(s[i+2:], s[i+1:])
			s[i+1] = b
			return s
		}

	}
	return s
}
