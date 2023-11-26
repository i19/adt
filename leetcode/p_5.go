package leetcode

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring
// 动态规划
//
//	func longestPalindrome(s string) string {
//		n := len(s)
//		if n <= 1 {
//			return s
//		}
//
//		dp := make([][]bool, len(s))
//		for i := 0; i < n; i++ {
//			dp[i] = make([]bool, len(s))
//		}
//
//		// 长度为 1
//		for i := 0; i < n; i++ {
//			dp[i][i] = true
//		}
//		start := 0
//		maxL := 1
//		// 长度为 2
//		for i := 0; i < n-1; i++ {
//			if s[i] == s[i+1] {
//				dp[i][i+1] = true
//				start = i
//				maxL = 2
//			}
//
//		}
//
//		// 检查长度大于2的子串
//		for l := 3; l <= n; l++ {
//			for i := 0; i < n-l+1; i++ {
//				j := i + l - 1
//				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
//				if dp[i][j] {
//					start = i
//					maxL = l
//				}
//			}
//		}
//
//		return s[start : start+maxL]
//	}
//
// 中心扩展
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandFromCenter(s, i, i)
		l2, r2 := expandFromCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}
func expandFromCenter(s string, left, right int) (int, int) {
	println("in ", left, right)
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}
