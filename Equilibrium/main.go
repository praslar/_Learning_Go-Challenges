package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{2, 7, 5},
		[]int{3, 1, 1},
		[]int{2, 1, -7},
		[]int{0, 2, 1},
		[]int{1, 6, 8},
	}
	fmt.Println(Solution(matrix))
}

func Solution(matrix [][]int) int {
	sumRow := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		sumR := 0
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
			sumR += matrix[i][j]
		}
		sumRow = append(sumRow, sumR)
		fmt.Println()
	}
	sumCol := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		sumC := 0
		for j := 0; j < len(matrix); j++ {
			fmt.Println(j, i)
			sumC += matrix[j][i]
		}
		sumCol = append(sumCol, sumC)
		fmt.Println()
	}
	fmt.Println("row", sumRow)
	fmt.Println("col", sumCol)
	//find possible N
	n := make([]int, 0)
	for i := 0; i < len(sumRow); i++ {
		//up
		sumUp := 0
		for j := i - 1; j >= 0; j-- {
			sumUp += sumRow[j]
		}
		//down
		sumDown := 0
		for k := i + 1; k < len(sumRow); k++ {
			sumDown += sumRow[k]
		}
		if sumDown == sumUp {
			n = append(n, i)
		}
	}
	//find possible Q
	q := make([]int, 0)
	for i := 0; i < len(sumCol); i++ {
		//up
		sumLeft := 0
		for j := i - 1; j >= 0; j-- {
			sumLeft += sumCol[j]
		}
		//down
		sumRight := 0
		for k := i + 1; k < len(sumCol); k++ {
			sumRight += sumCol[k]
		}
		if sumRight == sumLeft {
			q = append(q, i)
		}
	}

	rs := len(n) * len(q)
	fmt.Println(n)
	return rs
}
