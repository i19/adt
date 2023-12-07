package leetcode

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 1},
	}
	target := 2
	fmt.Println(searchMatrix(matrix, target))
}
