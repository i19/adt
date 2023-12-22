package leetcode

// 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position
func searchInsert(nums []int, target int) int {
	var helper func(left, right int) int
	helper = func(left, right int) int {
		if left > right {
			return left
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return helper(left, mid-1)
		} else {
			return helper(mid+1, right)
		}
	}
	return helper(0, len(nums)-1)
}
