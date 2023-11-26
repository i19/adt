package leetcode

// 516. 最长回文子序列
// https://leetcode.cn/problems/longest-palindromic-subsequence
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j+1])
			}
		}
	}
	return dp[n-1][0]
}
