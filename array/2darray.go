package main

import "fmt"

func TwoDimesionalArray() {
	var board [8][8]string
	board[0][0] = "r"
	board[0][1] = "n"

	for column := range board {
		board[1][column] = "p"
	}

	fmt.Println(board)
}

