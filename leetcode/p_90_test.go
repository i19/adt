package leetcode

import (
	"fmt"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	n := []int{4, 1, 4}
	for _, v := range subsetsWithDup(n) {
		fmt.Println(v)
	}
}
