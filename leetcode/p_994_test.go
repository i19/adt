package leetcode

import "testing"

func Test_orangesRotting(t *testing.T) {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	println(orangesRotting(grid))
}
