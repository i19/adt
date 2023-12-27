package leetcode

// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/

func robII(nums []int) int {
	helper := func(start, end int) int {
		f0, f1 := 0, 0
		for i := start; i < end; i++ {
			f0, f1 = f1, max(f0+nums[i], f1)
		}
		return f1
	}
	return max(
		nums[0]+helper(2, len(nums)-1),
		helper(1, len(nums)),
	)
}
