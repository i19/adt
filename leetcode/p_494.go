package leetcode

// 494. 目标和
// https://leetcode.cn/problems/target-sum/
// 思路
//
//	令
//		 s = sum(nums)
//		 p 为 nums 中添加 + 的所有数之和
//	那么
//		负数之和就是 s-p
//		p - (s-p) = target
//		2p = s + target
//		p = (s+target)/2
//	问题就变成了，在 nums 中寻找 和为 (s+target)/2 的 +数 的组成方案的数量 (选或者不选的问题，那就是 01 背包了)
//
//	dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-nums[i])
//
// s + target 是要 >= 0 的，同时也必须是 偶数
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s+target < 0 {
		return 0
	}
	if (s+target)%2 == 1 {
		return 0
	}

	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			} else {
				return 0
			}
		}
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}
	return dfs(len(nums)-1, (s+target)/2)
}

func findTargetSumWaysDPI(nums []int, target int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	target += s
	if target < 0 || target%2 == 1 {
		return 0
	}

	target /= 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1 // 对  应递归中的 i==0 && c==0
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] + dp[i][j-nums[i]]
			}
		}
	}
	return dp[len(nums)][target]
}

func findTargetSumWaysDPII(nums []int, target int) int {
	for _, v := range nums {
		target += v
	}
	if target < 0 || target%2 == 1 {
		return 0
	}

	target /= 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
}
