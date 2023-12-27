package leetcode

// 216. 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/
func combinationSum3(k int, n int) (res [][]int) {
	path := make([]int, k)
	var dfs func(pos, start, remain int)
	dfs = func(pos, start, remain int) {
		if pos == k && remain == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if pos == k {
			return
		}

		for i := start; i < 10 && i <= remain; i++ {
			path[pos] = i
			dfs(pos+1, i+1, remain-i)
		}
	}
	dfs(0, 1, n)
	return
}
