package own

// nums 为升序数组
func binarySearch(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return -1
}

// nums 为有序数组，将 target 插入到数组中保持有序特性，返回插入位置
// 如果 nums 中存在 target，则需要返回左侧的位置
func binaryInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在区间 [i, m-1] 中
			j = m - 1
		} else {
			// 首个小于 target 的元素在区间 [i, m-1] 中
			j = m - 1
		}
	}
	return i
}

// 复用查找左边界实现查找右边界
func binaryFindLeft(nums []int, target int) int {
	i := binaryInsert(nums, target)
	// 未找到 target ，返回 -1
	if i == len(nums) || nums[i] != target {
		return -1
	}
	// 找到 target ，返回索引 i
	return i
}

// 复用查找左边界实现查找右边界
func binaryFindRight(nums []int, target int) int {
	i := binaryFindLeft(nums, target+1)
	j := i - 1
	if j < 0 || nums[j] != target {
		return -1
	}
	return j
}
