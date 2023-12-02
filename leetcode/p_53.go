package leetcode

// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		if r < dp[i] {
			r = dp[i]
		}
	}
	return r
}
