package main

import "fmt"

var results map[string]int

func main() {
	fmt.Println(Solution([]int{1, 3, 2, 3, 4}))
}

func Solution(A []int) int {
	cut1, cut2, cut3 := getIndeces(A)
	rs := 0
	if cut1 == cut2 && cut2 == cut3 {
		return 0
	}
	for i := 0; i < 3; i++ {
		cut := make([]int, len(A))
		copy(cut, A)
		i1, i2, i3 := getIndeces(removeIndex(cut, cut1+i))
		if i1 == i2 && i2 == i3 {
			rs++
		}
	}
	if rs == 0 {
		return -1
	}

	return rs
}

func getIndeces(A []int) (int, int, int) {
	for i := 0; i < len(A)-2; i++ {
		if !isAsc(A[i], A[i+1], A[i+2]) {
			return i, i + 1, i + 2
		}
	}
	return 0, 0, 0
}

func isAsc(a int, b int, c int) bool {
	return ((a - b) * (b - c)) < 0
}

func removeIndex(a []int, i int) []int {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = -1
	a = a[:len(a)-1]
	return a
}
