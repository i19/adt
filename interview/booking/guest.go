package booking

import (
	"math"
	"sort"
)

// cap[i] 代表房间 i 的可承载客人数
// pirce[i] 代表方将 i 的价格
// total 代表客人总量
// 求能满足客人总量的最低价格
// 错误思路一，按照平均价格进行升序排序，然后从最低开始算，
// cap :=   []int{1,  2,  3}
// price := []int{10, 18, 24}
// count := 1
// 预期是 10， 但是会得到 24。
func q1(cap, price []int, total int) int {
	c := 0
	per := make([]int, len(cap))
	for i, v := range cap {
		per[i] = i
		c += v
	}
	if c < total {
		return -1
	}

	sort.Slice(per, func(i, j int) bool {
		return float32(price[i])/float32(cap[i]) < float32(price[j])/float32(cap[j])
	})

	ans := 0
	for i := 0; i < len(cap) && total > 0; i++ {
		ans += price[per[i]]
		total -= cap[per[i]]
	}
	return ans

}

// 暴力解决一切。用回溯
func q2(cap, price []int, total int) int {
	n := len(cap)
	ans := math.MaxInt32
	var dfs func(i, cost, remain int)
	dfs = func(i, cost, remain int) {
		if i >= n {
			if remain <= 0 {
				ans = min(ans, cost)
			}
			return
		}

		dfs(i+1, cost+price[i], remain-cap[i])
		dfs(i+1, cost, remain)
	}
	dfs(0, 0, total)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func q2Copy(cap, price []int, total int) int {
	n := len(cap)
	ans := math.MaxInt32
	var dfs func(i, cost, remain int)
	dfs = func(i, cost, remain int) {
		if remain <= 0 {
			ans = min(ans, cost)
			return
		}
		if i >= n {
			return
		}
		dfs(i+1, cost, remain)
		dfs(i+1, cost+price[i], remain-cap[i])
	}
	dfs(0, 0, total)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 回溯改成 DP
func q3(cap, price []int, total int) int {
	n := len(cap)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, total+1)
		for j := 1; j < total+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= total; j++ {
			if cap[i-1] >= j {
				dp[i][j] = min(dp[i-1][j], price[i-1])
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-cap[i-1]]+price[i-1])
			}
		}
	}
	if dp[n][total] == math.MaxInt32 {
		return -1
	} else {
		return dp[n][total]
	}
}

// DP 优化空间复杂度
func q4(cap, price []int, total int) int {
	dp := make([]int, total+1)
	for j := 1; j < total+1; j++ {
		dp[j] = math.MaxInt32
	}

	for i := 1; i <= len(cap); i++ {
		for j := total; j >= 1; j-- {
			if cap[i-1] >= j {
				dp[j] = min(dp[j], price[i-1])
			} else {

				dp[j] = min(dp[j], dp[j-cap[i-1]]+price[i-1])
			}
		}
	}
	if dp[total] == math.MaxInt32 {
		return -1
	} else {
		return dp[total]
	}
}
