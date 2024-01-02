package leetcode

import (
	"fmt"
	"testing"
)

func Test_highestRankedKItems(t *testing.T) {
	grid := [][]int{
		{1, 2, 0, 1},
		{1, 3, 0, 1},
		{0, 2, 5, 1},
	}
	price := []int{2, 5}
	start := []int{0, 0}
	k := 3
	fmt.Println(highestRankedKItems(grid, price, start, k))
}
