package leetcode

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	input := []int{1, 2, 3}

	for _, v := range permute(input) {
		fmt.Println(v)
	}
}
