package leetcode

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := &[]int{}
	var dfs func(node *TreeNode, remain int)
	dfs = func(node *TreeNode, remain int) {
		if node == nil {
			return
		}
		*path = append(*path, node.Val)
		defer func() { *path = (*path)[:len(*path)-1] }()
		if node.Val == remain && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int{}, (*path)...))
			return
		}
		dfs(node.Left, remain-node.Val)
		dfs(node.Right, remain-node.Val)
	}
	dfs(root, targetSum)
	return
}
