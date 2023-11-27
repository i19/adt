package leetcode

import (
	"fmt"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	for _, row := range board {
		for _, v := range row {
			fmt.Printf("%c ", v)
		}
		println()
	}
	println("========")

	// solveSudoku(board)
	// solveSudokuII(board)
	solveSudokuIII(board)
	for _, row := range board {
		for _, v := range row {
			fmt.Printf("%c ", v)
		}
		println()
	}
}
