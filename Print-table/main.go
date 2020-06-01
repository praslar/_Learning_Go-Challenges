package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := []int{4}
	Solution(number, 4)
}

func Solution(arr []int, col int) {

	//init table value
	row := len(arr) / col
	lastRow := col
	maxWidth := 0
	if len(arr)%col != 0 {
		row = (len(arr) / col) + 1
		lastRow = len(arr) % col
	}
	for _, number := range arr {
		if number > maxWidth {
			maxWidth = number
		}
	}
	maxWidth = len(strconv.Itoa(maxWidth))

	//draw
	numberDrawed := 0
	for i := 0; i < row; i++ {
		drawLine(maxWidth, col, len(arr))
		for j := 0; j < col && numberDrawed < len(arr); j++ {
			drawCell(maxWidth, col, arr[numberDrawed])
			numberDrawed++
		}
		fmt.Println("|")
	}
	drawLine(maxWidth, lastRow, len(arr))
}

func drawLine(maxWidth int, col int, lenArr int) {
	for i := 0; i < col && i < lenArr; i++ {
		fmt.Print("+")
		for i := 0; i < maxWidth; i++ {
			fmt.Print("-")
		}
	}
	fmt.Println("+")
}

func drawCell(maxWidth int, col int, number int) {
	numberLen := len(strconv.Itoa(number))
	fmt.Print("|")
	for i := 0; i < maxWidth-numberLen; i++ {
		fmt.Print(" ")
	}
	fmt.Print(number)
}
