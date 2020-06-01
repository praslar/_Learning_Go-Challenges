package main

import (
	"fmt"
	"sync"
)

func main() {
	b := []string{
		"..X....",
		"....X..",
		"....X..",
		".X..X..",
		"..X.X.X",
		"..X.X..",
		"...O...",
	}
	fmt.Println(Solution(b))
}

//Solution jafar
func Solution(b []string) int {

	// fan-in fan-out
	wg := &sync.WaitGroup{}
	stepChan := make(chan int)
	//init board
	board, jafarX, jafarY := initBoard(b)

	wg.Add(1)
	go move(board, jafarX, jafarY, 0, stepChan, wg)

	go func() {
		wg.Wait()
		close(stepChan)
	}()

	//lay gia tri cuoi cung trong close chanel
	data := 0
	for val := range stepChan {
		data = val
	}
	return data
}

func move(board [][]rune, x int, y int, step int, stepChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//left
	if x > 1 && y > 1 && board[x-1][y-1] == 'X' && board[x-2][y-2] == '.' {
		wg.Add(1)
		go move(board, x-2, y-2, step+1, stepChan, wg)
	}
	//right
	if x > 1 && y < (len(board)-2) && board[x-1][y+1] == 'X' && board[x-2][y+2] == '.' {
		wg.Add(1)
		go move(board, x-2, y+2, step+1, stepChan, wg)
	}
	stepChan <- step
}

func initBoard(b []string) ([][]rune, int, int) {
	boardSize := len(b)
	board := make([][]rune, boardSize)
	jarfarX, jafarY := 0, 0

	for i := 0; i < boardSize; i++ {
		board[i] = make([]rune, boardSize)
	}

	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			w := []rune(b[x])
			board[x][y] = w[y]
		}
	}

	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			//Get Jafar current location
			if board[x][y] == 'O' {
				jarfarX = x
				jafarY = y
			}
		}
	}
	return board, jarfarX, jafarY
}
