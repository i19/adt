package leetcode

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	input := []int{1, 1, 3, 1}

	for _, v := range permuteUnique(input) {
		fmt.Println(v)
	}
}
