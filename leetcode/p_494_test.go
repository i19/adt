package leetcode

import (
	"fmt"
	"testing"
)

func Test_findTargetSumWaysDP(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
	fmt.Println(findTargetSumWaysDPI(nums, target))
	fmt.Println(findTargetSumWaysDPII(nums, target))
}
