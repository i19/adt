package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	// nums := []int{0, 1, 0, 3, 2, 3}
	nums := []int{4, 10, 4, 3, 8, 9}
	println(lengthOfLIS(nums))
	println(lengthOfLISI(nums))
}
