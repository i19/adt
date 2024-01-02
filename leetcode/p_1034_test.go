package leetcode

import (
	"fmt"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	// grid := [][]int{
	// 	{1, 1},
	// 	{1, 2},
	// }
	// row := 0
	// col := 0
	// color := 3

	// grid := [][]int{
	// 	{1, 2, 2},
	// 	{2, 3, 2},
	// }
	// row, col, color := 0, 1, 3

	// grid := [][]int{
	// 	{1, 1, 1},
	// 	{1, 1, 1},
	// 	{1, 1, 1},
	// }
	// row, col, color := 1, 1, 2
	grid := [][]int{
		{1, 2, 1, 2, 1, 2},
		{2, 2, 2, 2, 1, 2},
		{1, 2, 2, 2, 1, 2},
	}
	row, col, color := 1, 3, 1
	fmt.Println(colorBorder(grid, row, col, color))
}
