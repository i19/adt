package leetcode

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUniqueHelper(state *[]int, res *[][]int, selected []bool, nums []int) {
	if len(*state) == len(nums) {
		*res = append(*res, append([]int{}, *state...))
	}

	used := make(map[int]bool)
	for i, v := range nums {
		if !selected[i] && !used[v] {
			*state = append(*state, v)
			selected[i] = true
			used[v] = true
			permuteUniqueHelper(state, res, selected, nums)
			selected[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}
func permuteUnique(nums []int) [][]int {
	var state []int
	var res [][]int
	selected := make([]bool, len(nums))
	permuteUniqueHelper(&state, &res, selected, nums)
	return res
}
