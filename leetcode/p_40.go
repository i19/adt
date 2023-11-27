package leetcode

import (
	"sort"
)

// 40. 组合总和 II
// https://leetcode.cn/problems/combination-sum-ii
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。

func combinationSum2Helper(start, remain int, candidates *[]int, state *[]int, res *[][]int) {
	if remain == 0 {
		*res = append(*res, append([]int{}, (*state)...))
		return
	}

	duplicated := make(map[int]bool)
	for i := start; i < len(*candidates); i++ {
		if (*candidates)[i] > remain {
			return
		}

		if !duplicated[(*candidates)[i]] {
			*state = append(*state, (*candidates)[i])
			duplicated[(*candidates)[i]] = true
			combinationSum2Helper(i+1, remain-(*candidates)[i], candidates, state, res)
			*state = (*state)[:len(*state)-1]
		}
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	// 配合大小剪枝，可以减少循环
	// 同时，如果想避免重复组合，引入的左边界机制，也需要 sorts 进行配合，因为这样可以去除最左侧元素，缩小搜索空间，从而避免重复组合
	sort.Ints(candidates)
	combinationSum2Helper(0, target, &candidates, &state, &res)
	return res
}
