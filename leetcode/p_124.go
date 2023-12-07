package leetcode

import "math"

// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

// 子树的最优不应是全局最优
// 所以可以考虑在求解子树最优的过程中，更新全局最优
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	// 路径包含该节点时，能返回的最大的『最大可用路径』
	// 这个 『最大可用路径』 和 『最大路径和』 是两个概念，因为返回的是路径，所以就不能再是 『拐弯』的的那种
	// 所以最后的 return 为 left 和 right 加了一个 max。来选择是返回该节点的左侧路径还是右侧路径
	// 同时在这个过程中，更新 『最大路径和』
	// 局部最优 vs 全局最优
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(helper(node.Left), 0)
		right := max(helper(node.Right), 0)
		// 引入该节点计算最大值，因为子树的已经和 maxSum 比较过了
		tsum := node.Val + left + right
		maxSum = max(maxSum, tsum)
		// 如果要返回给上一层，那么就只能是一条路径，而不是一棵树了
		return node.Val + max(left+right)
	}
	helper(root)
	return maxSum
}
