package leetcode

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	var findPos func(left, right int) int
	findPos = func(left, right int) int {
		if left > right {
			return -1
		}

		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return findPos(left, mid-1)
		} else {
			return findPos(mid+1, right)
		}
	}

	pos := findPos(0, len(nums)-1)
	if pos == -1 {
		return []int{-1, -1}
	}
	res := []int{pos, pos}

	for i := pos - 1; i >= 0 && nums[i] == nums[pos]; i-- {
		res[0] = i
	}
	for i := pos + 1; i < len(nums) && nums[i] == nums[pos]; i++ {
		res[1] = i
	}

	return res
}
