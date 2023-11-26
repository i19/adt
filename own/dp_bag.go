package own

import (
	"fmt"
	"math"
)

// 0-1 背包问题
// 给定 n 个物品，第 i 个物品的重量为 wg[i] 价值为 val[i] 和一个容量为 cap 的背包。
// 每个物品只能选择一次，问在不超过背包容量下能放入物品的最大价值。
// 解题思路 https://www.bilibili.com/video/BV1pY4y1J7na
// dp[i][j] 代表 包含物品 i 与 之前物品，在 容量为 j 时的最大价值
func dp01pack(wg []int, val []int, cap int) int {
	dp := make([][]int, len(wg)+1)
	for i := 0; i <= len(wg); i++ {
		dp[i] = make([]int, cap+1)
	}
	// 第一行和第一列默认是 0 ，不用再次初始化
	for item := 1; item <= len(wg); item++ {
		for tCap := 1; tCap <= cap; tCap++ {
			if tCap < wg[item-1] {
				// 无法选择当前物品，只能看之前的物品在当前容量下的最大收益
				dp[item][tCap] = dp[item-1][tCap]
			} else {
				dp[item][tCap] = max(
					// 不选择当前物品时，之前的物品在当前容量下的最大收益
					dp[item-1][tCap],
					// 之前物品（因为同一物品不能选择两次）在可用（剩余）容量下的最大收益 与 当前物品收益
					dp[item-1][tCap-wg[item-1]]+val[item-1],
				)
			}
		}
	}

	return dp[len(wg)][cap]
}

// 0-1 背包问题
// 针对空间做优化
// 观察代码发现， dp 中有用的一直是当前行和上一行
// 因此可以将 dp 缩减成 一维数组。
// 因为访问的上一行数据也都是当前位置的左侧，因此可以将数组从右向左更新（初始数据都是 0 ，所以是 OK 的）
// 并且这样做默认就继承了上一个数据（也就是不用体现 不选择当前物品的代码）
func dp01packII(wg []int, val []int, cap int) int {
	dp := make([]int, cap+1)
	for item := 1; item <= len(wg); item++ {
		for tCap := cap; tCap >= 1; tCap-- {
			if wg[item-1] <= tCap {
				dp[tCap] = max(
					dp[tCap],
					dp[tCap-wg[item-1]]+val[item-1],
				)
			}
		}
	}

	return dp[cap]
}

// 完全背包
// 与 0-1 背包类似，但是允许重复选择
func dpFullpack(wg []int, val []int, cap int) int {
	dp := make([][]int, len(wg)+1)
	for i := 0; i <= len(wg); i++ {
		dp[i] = make([]int, cap+1)
	}
	// 第一行和第一列默认是 0 ，不用再次初始化
	for item := 1; item <= len(wg); item++ {
		for tCap := 1; tCap <= cap; tCap++ {
			if tCap < wg[item-1] {
				dp[item][tCap] = dp[item-1][tCap]
			} else {
				dp[item][tCap] = max(
					dp[item-1][tCap],
					dp[item][tCap-wg[item-1]]+val[item-1],
				)
			}
		}
	}

	return dp[len(wg)][cap]
}

// 换零钱问题，也是一个 全背包
func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
		for j := 1; j <= amount; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(1+dp[i][j-coins[i-1]], dp[i-1][j])
			}
		}
	}
	if dp[len(coins)][amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[len(coins)][amount]
	}
}

// 换零钱问题，优化空间复杂度
// 逻辑很有意思
func coinChangeII(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		// 面值做起点，确保 dp[i-coin]可达
		for i := coin; i <= amount; i++ {
			// dp[i] 就是不使用该硬币的时候
			// dp[i-coin]+1 代表要新使用该面值(也就是之前可能用过了，也可能没用过)
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
		fmt.Printf("coin = %d, dp %v\n", coin, dp)
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// 换零钱问题，但是返回方案数量
func coinChangIII(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)

	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

// 换零钱问题，但是返回所有组合
func coinChangIIIIHelper(state *[]int, coins []int, leftBoundry, amount int, res *[][]int) {
	cAmount := 0
	for i, count := range *state {
		cAmount += coins[i] * count
	}
	if cAmount == amount {
		*res = append(*res, append([]int{}, *state...))
	}

	for i := leftBoundry; i < len(coins); i++ {
		if coins[i] <= amount-cAmount {
			(*state)[i]++
			// 这样将返回重复组合
			// coinChangIIIIHelper(state, coins, leftBoundry, amount, res)
			// 这样将不会返回重复组合，是正确结果
			coinChangIIIIHelper(state, coins, i, amount, res)
			(*state)[i]--
		}
	}
}
func coinChangIIII(coins []int, amount int) [][]int {
	state := make([]int, len(coins))
	res := make([][]int, 0)
	coinChangIIIIHelper(&state, coins, 0, amount, &res)
	return res
}
