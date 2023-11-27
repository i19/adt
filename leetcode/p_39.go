package leetcode

import "sort"

// 39. 组合总和
// https://leetcode.cn/problems/combination-sum
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
// 找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，
// 并以列表形式返回。你可以按 任意顺序 返回这些组合。
func combinationSumHelper(start, remain int, candidates *[]int, state *[]int, res *[][]int) {
	if remain == 0 {
		*res = append(*res, append([]int{}, (*state)...))
		return
	}

	for i := start; i < len(*candidates); i++ {
		if (*candidates)[i] > remain {
			return
		}

		*state = append(*state, (*candidates)[i])
		combinationSumHelper(i, remain-(*candidates)[i], candidates, state, res)
		*state = (*state)[:len(*state)-1]
	}
}
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	// 配合大小剪枝，可以减少循环
	sort.Ints(candidates)
	combinationSumHelper(0, target, &candidates, &state, &res)
	return res
}
