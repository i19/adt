package leetcode

import (
	"fmt"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	input := [][][]int{
		[][]int{
			{1},
		},
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
	}

	for i := range input {
		rotate(input[i])
		for _, v := range input[i] {
			fmt.Println(v)
		}
		println("========")

	}
}
