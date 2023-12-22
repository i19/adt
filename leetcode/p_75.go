package leetcode

// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors
func sortColors(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

	for i := p; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
}
