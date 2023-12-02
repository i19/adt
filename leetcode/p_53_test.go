package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums = []int{-1}
	// nums = []int{-1, 1, 4, -4, 5}
	// nums = []int{-1, -2}

	println(maxSubArray(nums))
}
