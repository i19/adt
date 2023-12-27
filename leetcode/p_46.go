package leetcode

// 46. 全排列
// https://leetcode.cn/problems/permutations
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permuteHelper(state *[]int, res *[][]int, selected []bool, nums []int) {
	if len(*state) == len(nums) {
		*res = append(*res, append([]int{}, *state...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !selected[i] {
			*state = append(*state, nums[i])
			selected[i] = true
			permuteHelper(state, res, selected, nums)
			*state = (*state)[:len(*state)-1]
			selected[i] = false
		}
	}
}
func permute(nums []int) [][]int {
	var state []int
	var res [][]int
	selected := make([]bool, len(nums))
	permuteHelper(&state, &res, selected, nums)
	return res
}

func permuteII(nums []int) (res [][]int) {
	path := make([]int, len(nums))
	selected := make([]bool, len(nums))
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, v := range nums {
			if !selected[i] {
				selected[i] = true
				path[pos] = v
				dfs(pos + 1)
				selected[i] = false
			}
		}
	}
	dfs(0)
	return
}
