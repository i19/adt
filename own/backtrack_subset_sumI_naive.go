package own

import "sort"

// 子集和问题

// case1: 无重复元素的情况
// 给定一个正整数数组 nums 和一个目标正整数 target ，
// 请找出所有可能的组合，使得组合中的元素和等于 target 。
// 给定数组无重复元素，每个元素可以被选取多次。请以列表形式返回这些组合，列表中不应包含重复组合。
// 剪枝1: 和 大于 target
// 剪枝2: 为了去除 重复组合，引入了 start ，使 choice 限制为 只能使用最初节点（含）以及之后的节点，从而避免后面的节点使用前面的节点，造成 『重复』
func ssnI(nums []int, target int) [][]int {
	state := []int{}
	res := [][]int{}
	start := 0

	var dfs func(start int, state *[]int, nums []int, target int, res *[][]int)
	dfs = func(start int, state *[]int, nums []int, target int, res *[][]int) {
		sum := 0
		for _, v := range *state {
			sum += v
		}

		if sum == target {
			// 结果保存
			v := append([]int{}, *state...)
			*res = append(*res, v)
		} else if sum > target {
			// 越界剪枝
			return
		}

		// 规定可用节点左边界，针对 重复组合 进行剪枝
		for i := start; i < len(nums); i++ {
			// 做选择、更新状态
			// (并且，这样可以将当前节点放入，也就是必须选择当前节点，从而避免一种情况：虽然给的范围很大，但是只从后面选择，这样还是会出现重复)
			*state = append(*state, nums[i])
			// 尝试求解
			dfs(i, state, nums, target, res)
			// 状态回溯
			*state = (*state)[:len(*state)-1]
		}
	}
	dfs(start, &state, nums, target, &res)

	return res
}

// case2: 有重复元素的情况
// 给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，
// 使得组合中的元素和等于 target 。给定数组可能包含重复元素，每个元素只可被选择一次。
// 请以列表形式返回这些组合，列表中不应包含重复组合。
// 剪枝1: 和 大于 target
// 剪枝2: 为了去除 重复组合，引入了 start ，使 choice 限制为 只能使用最初节点（含）以及之后的节点，从而避免后面的节点使用前面的节点，造成 『重复』
// 剪枝3: 为了去除重复元素，需要确保同一轮选择中，不能出现 duplicate
func ssnII(nums []int, target int) [][]int {
	state := []int{}
	res := [][]int{}
	start := 0
	sort.Ints(nums)

	var dfs func(start int, state *[]int, nums []int, remain int, res *[][]int)
	dfs = func(start int, state *[]int, nums []int, remain int, res *[][]int) {
		if remain == 0 {
			v := append([]int{}, *state...)
			*res = append(*res, v)
		}

		duplicate := make(map[int]struct{})
		// 规定可用节点左边界，针对 重复组合 以及越界 进行剪枝
		for i := start; i < len(nums) && nums[i] <= remain; i++ {
			if _, ok := duplicate[nums[i]]; !ok {
				duplicate[nums[i]] = struct{}{}
				// 做选择、更新状态
				*state = append(*state, nums[i])
				// 尝试求解
				dfs(i, state, nums, remain-nums[i], res)
				// 状态回溯
				*state = (*state)[:len(*state)-1]
			}
		}
	}
	dfs(start, &state, nums, target, &res)
	return res
}
