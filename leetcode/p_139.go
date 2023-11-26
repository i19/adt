package leetcode

// 139. 单词拆分
// https://leetcode.cn/problems/word-break
func wordBreak(s string, wordDict []string) bool {
	x := make(map[string]bool)
	for _, v := range wordDict {
		x[v] = true
	}
	// 多一个，方便设置初始条件，否则循环无法执行
	// 所以循环也是从 1 开始的，正好配合 前开后闭
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && x[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
