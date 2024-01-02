package own

// 01 背包
// 01 的意思是，要么选，要么不选，不能部分选
// 有 n 件物品和一个容量是 c 的背包。每件物品只能使用一次。
// 第 i 件物品的体积是 c[i]，价值是 v[i]
// 求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。输出最大价值。

// 回溯三问
// 当前操作?
//
//	枚举第i个物品选或不选:
//		不选，剩余容量不变;
//		选,  剩余容量减少 c[i]
//
// 子问题?
//
//	在剩余容量为 c 时, 从前 i-1 个物品中得到的最大价值和
//
// 下一个子问题?
//
//		分类讨论:
//			不选: 在剩余容量为 c 时从前 i-1 个物品中得到的最大价值和
//			选:  在剩余容量为 c-c[i]时,从前 i-1 个物品中得到的最大价值和
//	     得出方程 dfs(i,c) = max(dfs(i-1, c), dfs(i-1, c-c[i]) + v[i])
func zero_one_knapsack(cs, vs []int, maxC int) int {
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		// 没有物品可供选择
		if i < 0 {
			return 0
		}
		// 物品体积超了
		if cs[i] > c {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i-1, c-cs[i])+vs[i])
	}
	return dfs(len(cs)-1, maxC)
}

// dp 自顶向下
func zero_one_knapsack_memo(cs, vs []int, maxC int) int {
	memo := make(map[int]map[int]int)
	for i := 0; i < len(vs); i++ {
		memo[i] = make(map[int]int)
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		// 没有物品可供选择
		if i < 0 {
			return 0
		}
		if v, ok := memo[i][c]; ok {
			return v
		}
		// 物品体积超了
		if cs[i] > c {
			memo[i][c] = dfs(i-1, c)
			return memo[i][c]
		}
		memo[i][c] = max(memo[i][c], dfs(i-1, c-cs[i])+vs[i])
		return memo[i][c]
	}
	x := dfs(len(cs)-1, maxC)
	return x
}

// dp 自底向上
func zero_one_knapsack_DP(cs, vs []int, maxC int) int {
	memo := make(map[int]map[int]int)
	for i := 0; i <= len(vs); i++ {
		memo[i] = make(map[int]int)
	}
	memo[-1] = make(map[int]int)
	for i := 0; i < len(vs); i++ {
		for j := 1; j <= maxC; j++ {
			if j < cs[i] {
				memo[i][j] = memo[i-1][j]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i-1][j-cs[i]]+vs[i])
			}
		}
	}
	return memo[len(cs)-1][maxC]
}

// dp 自底向上 优化空间
// 注意这里对容量遍历是反过来的
// 因为访问的上一行数据也都是当前位置的左侧，因此需要将数组从右向左更新（初始数据都是 0 ，所以是 OK 的）
// 并且这样做默认就继承了上一个数据（也就是不用体现 不选择当前物品的代码）
func zero_one_knapsack_DP_space(cs, vs []int, maxC int) int {
	memo := make([]int, maxC+1)
	for i := 0; i < len(vs); i++ {
		for j := maxC; j >= 1; j-- {
			if j >= cs[i] {
				memo[j] = max(memo[j], memo[j-cs[i]]+vs[i])
			}
		}
	}
	return memo[maxC]
}

// 完全背包
// 有 n 种物品和一个容量是 c 的背包。每种物品不再限制使用一次。
// 第 i 种物品的体积是 c[i]，价值是 v[i]
func unbound_knapsack(cs, vs []int, maxC int) int {
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		// 没有物品可供选择
		if i < 0 {
			return 0
		}
		// 物品体积超了
		if cs[i] > c {
			return dfs(i-1, c)
		}
		// 相比 01 背包，只有这里变了，从 i-1 编程了 i（代表该物体仍旧可以选择）
		return max(dfs(i, c), dfs(i, c-cs[i])+vs[i])
	}
	return dfs(len(cs)-1, maxC)
}
