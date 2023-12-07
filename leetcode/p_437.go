package leetcode

// 437. 路径总和 III
// https://leetcode.cn/problems/path-sum-iii/description

// 深度优先、回溯、双递归
func pathSumI(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var rootSum func(node *TreeNode, targetSum int) int
	rootSum = func(node *TreeNode, targetSum int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val == targetSum {
			res++
		}
		res += rootSum(node.Left, targetSum-node.Val)
		res += rootSum(node.Right, targetSum-node.Val)
		return res
	}

	var res int
	res += rootSum(root, targetSum)
	res += pathSumI(root.Left, targetSum)
	res += pathSumI(root.Right, targetSum)
	return res
}

// 利用前缀和将上面的代码进行优化
func pathSumII(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return 0
	}

	preSum := map[int]int{0: 1}
	var helper func(node *TreeNode, cur int)
	helper = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		// 结果保存
		res += preSum[cur-targetSum]
		// 将前缀和 进行选择
		preSum[cur]++
		helper(node.Left, cur)
		helper(node.Right, cur)
		// 将前缀和 进行回溯
		preSum[cur]--
	}

	helper(root, 0)
	return
}
