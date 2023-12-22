package leetcode

// 153. 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 如果到了右区间，那么直接返回 nums[l] 就好，不需要 r 进行递进了
		if nums[l] < nums[r] {
			return nums[l]
		}

		// 否则还是需要两边进行靠近
		mid := (l + r) / 2

		if nums[mid] >= nums[l] && nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
