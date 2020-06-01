package main

import "fmt"

func main() {
	fmt.Println(Solution([][]int{
		{2, 1, 1, 4},
	}))
}

func Solution(A [][]int) int {
	col := len(A[0])
	row := len(A)

	//Max magic size
	maxSize := row
	if row > col {
		maxSize = col
	}

	for magicSize := maxSize; magicSize > 1; magicSize-- {
		for h := 0; h <= (row - magicSize); h++ {
			for w := 0; w <= (col - magicSize); w++ {
				if isMagic(A, h, w, magicSize) {
					return magicSize
				}
			}
		}

	}
	return 1
}

func isMagic(A [][]int, h, w, magicSize int) bool {
	for i := 0; i < magicSize; i++ {
		sumRow, sumCol, sumDgn := 0, 0, 0
		for j := 0; j < magicSize; j++ {
			sumRow += A[i+h][j+w]
			sumCol += A[j+h][i+w]
			sumDgn += A[j+h][j+w]
		}
		if sumRow != sumCol || sumRow != sumDgn {
			return false
		}
	}
	return true
}
