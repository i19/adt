package leetcode

// 77. 组合
// https://leetcode.cn/problems/combinations
func combine(n int, k int) [][]int {
	stats := []int{}
	res := [][]int{}
	var helper func(pos int)
	helper = func(pos int) {
		if len(stats) == k {
			res = append(res, append([]int{}, stats...))
		}
		for i := pos; i <= n; i++ {
			stats = append(stats, i)
			helper(i + 1)
			stats = stats[:len(stats)-1]
		}
	}
	helper(1)
	return res
}

// 简单优化下空间使用部分
func combineII(n int, k int) [][]int {
	stats := make([]int, 0, k)
	res := [][]int{}
	var helper func(pos int)
	helper = func(pos int) {
		if len(stats) == k {
			temp := make([]int, k)
			copy(temp, stats)
			res = append(res, temp)
		}
		for i := pos; i <= n; i++ {
			stats = append(stats, i)
			helper(i + 1)
			stats = stats[:len(stats)-1]
		}
	}
	helper(1)
	return res
}
