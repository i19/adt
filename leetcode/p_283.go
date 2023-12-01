package leetcode

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes
func moveZeroes(nums []int) {
	zeroCount := 0
	for i, v := range nums {
		if v == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = v
		}
	}
	for i := len(nums) - 1; zeroCount > 0; zeroCount-- {
		nums[i] = 0
		i--
	}
}
