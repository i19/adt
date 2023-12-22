package leetcode

import (
	"fmt"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 4}
	findDuplicate(nums)
	fmt.Println(nums)
}
