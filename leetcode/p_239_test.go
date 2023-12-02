package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindowp(t *testing.T) {
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{1, -1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}
