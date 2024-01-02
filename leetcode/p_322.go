package leetcode

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/
import "math"

// 回溯机制
func coinChange(coins []int, amount int) int {
	var dfs func(i, remain int) int
	dfs = func(i, remain int) int {
		if i < 0 {
			if remain == 0 {
				return 0
			}
			return math.MaxInt32
		}
		if remain < coins[i] {
			return dfs(i-1, remain)
		}
		// 选不选
		return min(dfs(i-1, remain), dfs(i, remain-coins[i])+1)
	}

	ans := dfs(len(coins)-1, amount)
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func coinChangeDPI(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 1; j <= amount; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	ans := dp[n][amount]
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func coinChangeDPII(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}
