package leetcode

// 2140. 解决智力问题
// https://leetcode.cn/problems/solving-questions-with-brainpower
// 反推
func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions)+1)
	for i := len(questions) - 1; i >= 0; i-- {
		dp[i] = max(
			dp[i+1],
			questions[i][0]+dp[min(len(questions), i+questions[i][1]+1)])
	}
	return int64(dp[0])
}
