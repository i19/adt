package introduction_to_algorithms

import "fmt"

// 动态规划相关
// 以 斐波那契 做示例，演示两种机制

// FibRaw 传统 Fib 算法
func FibRaw(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibRaw(n-1) + FibRaw(n-2)
	}
}

// FibBottom 自底向上（常见的 DP 写法）
func FibBottom(n int) int {
	c := []int{0, 1}
	if n < 2 {
		return c[n]
	}

	for j := 2; j <= n; j++ {
		c = append(c, c[j-2]+c[j-1])
	}
	return c[n]
}

// FibWithMemo 改造 FibRaw 成 （带有备忘录的）
// 逻辑是自顶向下，通过递归的形式，备忘录的实际写入仍旧是自底向上
func FibWithMemo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		memo := make([]int, n+1)
		memo[0], memo[1] = 0, 1
		return helper(n, memo)
	}
}
func helper(n int, memo []int) int {
	if n < 2 {
		return memo[n]
	}
	if memo[n] == 0 {
		memo[n] = helper(n-1, memo) + helper(n-2, memo)
		return memo[n]
	} else {
		return memo[n]
	}
}

// 最长子序列
func findLCS(X, Y string) string {
	m := len(X)
	n := len(Y)

	// Initialize the DP table
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Compute the lengths of LCS
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp)

	// Reconstruct the LCS
	// elseif 和 else 两个地方 i-- 和 j-- 的地方，是按照最大长度进行回溯
	lcs := ""
	i, j := m, n
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			lcs = string(X[i-1]) + lcs
			println(i, j, dp[i][j])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最长的单调递增子序列
// findIncreaseLCS 请给出 一个O(n^2)时间的算法，使之能找出一个 n 个数的序列中最长的单调递增子序列
// 单调递增的要求，导致该问题不适合使用 DP，因为不存在最优子结构  (这是一个错误的示例，这是我写的)
// 下面的 findLongestIncreasingSubsequence 是 chatGPT 写的一个正确答案，用到了 DP
// 我理解错误了 .... 对比着看，收益匪浅
func findIncreaseLCS(in []int) []int {
	r := make([]int, len(in))
	var rr []int
	rmax := 0
	for i := 1; i < len(in); i++ {
		n := 1
		subR := []int{in[i]}
		for j := i; j > 0; j-- {
			if in[j-1] <= in[j] {
				n++
				subR = append([]int{in[j-1]}, subR...)
			}
		}
		r[i] = n
		if rmax < n {
			rmax = n
			rr = subR
		}
	}
	fmt.Println(r)
	return rr
}

func findLongestIncreasingSubsequence(sequence []int) []int {
	n := len(sequence)

	// Initialize the dp array with 1s
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1 // Length of the longest increasing subsequence

	// Compute the longest increasing subsequence
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if sequence[i] > sequence[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	fmt.Println(dp)

	// Construct the longest increasing subsequence
	// 这个利用变化，来存储值的思路是真牛逼（变化了，说明这里的值有用）
	// ！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
	subsequence := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == maxLength {
			subsequence = append([]int{sequence[i]}, subsequence...)
			maxLength--
		}
	}
	// ！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！

	return subsequence
}
