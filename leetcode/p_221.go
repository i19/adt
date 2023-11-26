package leetcode

// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square/?envType=study-plan-v2&envId=dynamic-programming
// 思路：寻找最长的边，而不是面积计算
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxRes := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxRes = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxRes {
					maxRes = dp[i][j]
				}
			}
		}
	}
	return maxRes * maxRes
}
