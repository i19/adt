package leetcode

import (
	"fmt"
	"testing"
)

func Test_pathWithObstacles(t *testing.T) {
	board := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	// board := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	// board := [][]int{{0}}
	fmt.Println(pathWithObstaclesI(board))
}
