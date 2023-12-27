package leetcode

// 131. 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/
// 示例 1：
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]

// 思路
// 将 "aab", 看成是 "a,a,b"
// 然后对每个 ',' 做 选还是不选 的处理
// 从而看成回溯的 子集型回溯

func isEchoString(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 站在输入角度，对元素的动作有两种，选择或者不选
func partitionI(s string) (res [][]string) {
	var path []string
	var dfs func(i, start int)
	dfs = func(i, start int) {
		if i == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		// 不选择该元素
		if i < len(s)-1 {
			dfs(i+1, start)
		}

		// 选择该元素
		if isEchoString(s[start : i+1]) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]

		}
	}
	dfs(0, 0)
	return
}

// 站在答案角度
func partitionII(s string) (res [][]string) {
	var path []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for j := i; j < len(s); j++ {
			t := s[i : j+1]
			if isEchoString(t) {
				path = append(path, t)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
