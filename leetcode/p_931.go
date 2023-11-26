package leetcode

// 931. 下降路径最小和
// https://leetcode.cn/problems/minimum-falling-path-sum/?envType=study-plan-v2&envId=dynamic-programming

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix) == 1 {
		return min(matrix[0]...)
	}
	if len(matrix[0]) == 1 {
		return sum(matrix[0])
	}

	rowM := len(matrix)
	colM := len(matrix[0])

	dp := make([][]int, rowM)
	for i := 0; i < rowM; i++ {
		dp[i] = make([]int, colM)
	}
	for row := 0; row < rowM; row++ {
		for col := 0; col < colM; col++ {
			if row == 0 {
				dp[row][col] = matrix[row][col]
			} else if col == 0 {
				dp[row][col] = matrix[row][col] + min(dp[row-1][0], dp[row-1][1])
			} else if col == colM-1 {
				dp[row][col] = matrix[row][col] + min(dp[row-1][col-1], dp[row-1][col])
			} else {
				dp[row][col] = matrix[row][col] + min(dp[row-1][col-1], dp[row-1][col], dp[row-1][col+1])
			}
		}
	}
	return min(dp[len(matrix)-1]...)
}

func sum(args []int) int {
	v := 0
	for _, x := range args {
		v += x
	}
	return v
}
