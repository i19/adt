package leetcode

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	// nums := []int{2, 1, 5, 6, 2, 3}
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums := []int{2, 4}
	println(largestRectangleArea(nums))
	println(largestRectangleAreaII(nums))
}
