package leetcode

import (
	"fmt"
	"testing"
)

func Test_findTheCity(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 2},
		{0, 4, 8},
		{1, 2, 3},
		{1, 4, 2},
		{2, 3, 1},
		{3, 4, 1},
	}
	distanceThreshold := 2

	fmt.Println(findTheCity(n, edges, distanceThreshold))
}
