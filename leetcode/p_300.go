package leetcode

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence
// 尝试使用单调栈
// （单调栈不行，会找第一个，但是 子序列是可以 跳着组成）

func lengthOfLISI(nums []int) int {
	prev := make([]int, len(nums))
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			prev[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	res := 0
	for start := 0; start < (len(prev) - res); start++ {
		tres := 1
		for next := start; prev[next] != 0 && prev[next] < len(nums); next = prev[next] {
			tres++
		}
		res = max(res, tres)
	}
	return res
}

// 使用动态规划
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				ans = max(ans, dp[i])
			}
		}

	}
	return ans
}
