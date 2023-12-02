package leetcode

import (
	"fmt"
	"testing"
)

func Test_mergeI(t *testing.T) {
	nums := [][]int{{1, 3}, {1, 2}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(mergeI(nums))
	fmt.Println(mergeII(nums))
	fmt.Println(mergeIII(nums))
}
