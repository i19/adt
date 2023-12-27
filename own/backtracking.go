package own

// 回溯算法 backtracking algorithm 是一种通过穷举来解决问题的方法，
// 它的核心思想是从一个初始状态出发，暴力搜索所有可能的解决方案，当遇到正确的解则将其记录，
// 直到找到解或者尝试了所有可能的选择都无法找到解为止。
// 回溯算法通常采用“深度优先搜索”来遍历解空间。在二叉树章节中，
// 我们提到前序、中序和后序遍历都属于深度优先搜索

// “尝试”与“回退”
// 之所以称之为回溯算法，是因为该算法在搜索解空间时会采用“尝试”与“回退”的策略
// 在二叉树中搜索所有值为 7 的节点，请返回根节点到这些节点的路径
func preOrderII(root *treeNode, res *[][]int, path *[]int) {
	if root == nil {
		return
	}

	// 尝试
	*path = append(*path, root.Val)
	if root.Val == 7 {
		// 记录解
		*res = append(*res, *path)
	}

	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

// 剪枝
// 复杂的回溯问题通常包含一个或多个约束条件，约束条件通常可用于“剪枝”。
// 在二叉树中搜索所有值为 7 的节点，请返回根节点到这些节点的路径，并要求路径中不包含值为 3 的节点。
func preOrderIII(root *treeNode, res *[][]int, path *[]int) {
	// 剪枝
	if root == nil || root.Val == 3 {
		return
	}

	// 尝试
	*path = append(*path, root.Val)
	if root.Val == 7 {
		// 记录解
		*res = append(*res, *path)
	}

	preOrderIII(root.Left, res, path)
	preOrderIII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

// 回溯算法 框架
// 以 preOrderIII 的问题为例子
// 当前的状态是否为解（当前路径是否到达了 7）
func isSolution(state *[]*treeNode) bool {
	return len(*state) != 0 && (*state)[len(*state)-1].Val == 7
}

// 将当前状态记录为结果（将当前的路径信息进行保存）
func recordSolution(state *[]*treeNode, res *[][]int) {
	var solution []int
	for _, v := range *state {
		solution = append(solution, v.Val)
	}
	*res = append(*res, solution)
}

// 用于『剪枝』
func isValid(state *[]*treeNode, choice *treeNode) bool {
	return choice != nil && choice.Val != 3
}

// 做选择，同时将选择记录成为状态（路径的更新，尝试）
func makeChoice(state *[]*treeNode, choice *treeNode) {
	*state = append(*state, choice)
}

// 状态回退（路径信息往回走）
func undoChoice(state *[]*treeNode, choice *treeNode) {
	*state = (*state)[:len(*state)-1]
}

// 回溯的框架结构
// 需要定义的就是 状态 和 当前状态的下一步选择
// 在该问题里
//
//	state 就是当前的路径信息
//	choices 则是在当前已有的路径下一步的选择（往左走，往右走）
func backtrackIII(state *[]*treeNode, choices *[]*treeNode, res *[][]int) {
	if isSolution(state) {
		// 结果记录
		recordSolution(state, res)
	}

	for _, choice := range *choices {
		// 剪枝
		if isValid(state, choice) {
			// 做选择
			makeChoice(state, choice)
			// 下一步的尝试
			nextChoice := []*treeNode{choice.Left, choice.Right}
			backtrackIII(state, &nextChoice, res)
			// 状态回溯
			undoChoice(state, choice)
		}
	}
}

func permutationsIHelper(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 结果保存
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}

	for i := 0; i < len(*choices); i++ {
		// 剪枝（不可使用已使用过的元素）
		if !(*selected)[i] {
			// 做选择，状态更新
			(*selected)[i] = true
			*state = append(*state, (*choices)[i])
			permutationsIHelper(state, choices, selected, res)
			// 状态回溯
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

// 全排列问题
// 全排列 输入一个整数数组，数组中不包含重复元素，返回所有可能的排列。
func permutationsI(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	// selected 其实也是 状态 一部分，用于剪枝
	selected := make([]bool, len(nums))
	permutationsIHelper(&state, &nums, &selected, &res)
	return res
}

// 全排列的框架实现
func permutationsIIIsSolution(state *[]int, nums *[]int) bool {
	return len(*state) == len(*nums)
}

func permutationsIIIsValid(selected *[]bool, choice int) bool {
	return !(*selected)[choice]
}

func permutationsIIMakeChoice(state *[]int, selected *[]bool, choice int, nums *[]int) {
	*state = append(*state, (*nums)[choice])
	(*selected)[choice] = true
}
func permutationsIIUndoChoice(state *[]int, selected *[]bool, choice int) {
	*state = (*state)[:len(*state)-1]
	(*selected)[choice] = false
}
func permutationsIIRecordSolution(state *[]int, res *[][]int) {
	v := append([]int{}, *state...)
	*res = append(*res, v)
}

func permutationsIIHelper(state *[]int, res *[][]int, selected *[]bool, nums *[]int) {
	// 保存结果
	if permutationsIIIsSolution(state, nums) {
		permutationsIIRecordSolution(state, res)
	}

	// 用下标当选择
	for choice := range *nums {
		// 剪枝
		if permutationsIIIsValid(selected, choice) {
			// 做选择（修改状态，状态包含路径和用以剪枝的 selected 更新）
			permutationsIIMakeChoice(state, selected, choice, nums)
			// 尝试（对新状态进行判断、扩展。） 回溯的本质是深度优先搜索
			permutationsIIHelper(state, res, selected, nums)
			// 状态回溯
			permutationsIIUndoChoice(state, selected, choice)
		}
	}
}
func permutationsII(nums []int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	selected := make([]bool, len(nums))
	permutationsIIHelper(&state, &res, &selected, &nums)
	return res
}

// 全排列，输入一个整数数组，数组中可能包含重复元素，返回所有不重复的排列。
func permutationsIII(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	// selected 用于重复选择剪枝
	// 		整个搜索过程中只有一个 selected 。它记录的是当前状态中包含哪些元素，
	// 		作用是防止 choices 中的任一元素在 state 中重复出现。
	selected := make([]bool, len(nums))
	permutationsIIIHelper(&state, &nums, &selected, &res)
	return res
}

func permutationsIIIHelper(state, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	// duplicated 放在循环之外，同样是为了剪枝
	// 在同一轮内做选择时，去除重复选择
	// duplicated相等元素剪枝
	// 		每轮选择（即每个调用的 backtrack 函数）都包含一个 duplicated 。
	// 		它记录的是在本轮遍历（即 for 循环）中哪些元素已被选择过，作用是保证相等的元素只被选择一次。
	duplicated := make(map[int]struct{}, 0)
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素 且 不允许重复选择相等元素
		if _, ok := duplicated[choice]; !ok && !(*selected)[i] {
			// 尝试：做出选择，更新状态
			// 记录选择过的元素值
			duplicated[choice] = struct{}{}
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			permutationsIIIHelper(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}
