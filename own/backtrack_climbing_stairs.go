package own

// 爬楼梯
// 给定一个共有 n 阶的楼梯，你每步可以上 1 阶或者 2 阶，请问有多少种方案可以爬到楼顶。
// 该问题其实是 DP 的常见题，下面用 回溯算法 进行解决
func climbingStairsBacktrackHelper(state int, target int, choices []int, res *int) {
	// 存储结果
	if state == target {
		(*res)++
		return
	}

	// 原始方法
	for _, choice := range choices {
		// 做选择
		state += choice
		// 剪枝
		if state > target {
			break
		}
		// 解题
		climbingStairsBacktrackHelper(state, target, choices, res)
		// 回溯
		state -= choice
	}

	// 变体 1
	// for _, choice := range choices {
	// 	if state > target {
	// 		break
	// 	}
	// 	climbingStairsBacktrackHelper(state+choice, target, choices, res)
	// }

	// 变体 2
	// if state > target {
	// 	return
	// }
	// climbingStairsBacktrackHelper(state+1, target, res)
	// climbingStairsBacktrackHelper(state+2, target, res)

}
func climbingStairsBacktrack(n int) int {
	choice := []int{1, 2}
	state := 0
	res := 0
	climbingStairsBacktrackHelper(state, n, choice, &res)
	return res
}

// 根据 f(i) = f(i-1) + f(i-2) 使用递归方式
// 这意味着在爬楼梯问题中，**各个子问题之间存在递推关系，原问题的解可以由子问题的解构建得来**
// 这种递归意味着中间有很多重复计算，DP 的意义就是通过某种的形式保存这种中间结果，避免重复计算
func climbingStairsDFS(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbingStairsDFS(n-1) + climbingStairsDFS(n-2)
}

// 下面就是带有 memo 的 DP 实现
func climbingStairsDPMemo(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 观察 climbingStairsDPMemo 会发现计算 dp[i] 的时候一直用的是 i-1 和 i-2 的value，并没有使用更向前的值
// 因此可以将空间占用从 n -> 1
// 该种方法适合很多 DP 的 memo 修改
func climbingStairsDPWithoutMemo(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	p1, p2, r := 1, 2, 0
	for i := 3; i <= n; i++ {
		r = p1 + p2
		p1 = p2
		p2 = r
	}
	return r
}

// DP 的无后效性
// 无后效性是动态规划能够有效解决问题的重要特性之一
// 定义为：给定一个确定的状态，它的未来发展只与当前状态有关，而与当前状态过去所经历过的所有状态无关。
// 现在为爬楼梯问题增加一个约束
// 给定一个共有 n 阶的楼梯，你每步可以上 1 阶或者 2 阶，但不能连续两轮跳
// 阶，请问有多少种方案可以爬到楼顶。
// 此时，之前的 f(i) = f(i-1) + f(i-2) 状态转移方程失效了
// 这个特定的问题可以通过扩展 状态 来解决，是问题重新满足 无后效性 可以继续使用 DP 解决
// 但是许多更复杂的组合优化问题都不满足无后效性，此时会使用其他算法。比如 启发式搜索、遗传算法、强化学习 ...
func climbingStairsConstraintDP(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([][3]int, n+1)
	// 初始状态：预设最小子问题的解
	// dp[i][j] 代表处在第 i 阶，并且上一轮跳了 j 阶
	dp[1][1] = 1
	dp[1][2] = 0
	dp[2][1] = 0
	dp[2][2] = 1
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i][1] = dp[i-1][2]
		dp[i][2] = dp[i-2][1] + dp[i-2][2]
	}
	return dp[n][1] + dp[n][2]
}
