package leetcode

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{5, 1, 1}
	// nums := []int{2, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
