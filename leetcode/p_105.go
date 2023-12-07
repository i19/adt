package leetcode

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal

// layer    1 2 3 4 5 6 7
// preorder 1 2 4 5 3 6 7
// inorder  4 2 5 1 6 3 7

// preorder 中->左->右
// inorder 左->中->右
// 利用 preorder 获取子树根节点的值，并通过递归的方式将该节点进行返回
// 利用 inorder 确定子树左右子树范围，协助定位 根节点位置
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var helper func(rootPrePos, inLeft, inRight int) *TreeNode
	helper = func(rootPrePos, inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}
		root := &TreeNode{Val: preorder[rootPrePos]}
		inRoot := inorderMap[preorder[rootPrePos]]

		root.Left = helper(rootPrePos+1, inLeft, inRoot-1)
		root.Right = helper(rootPrePos+1+inRoot-inLeft, inRoot+1, inRight)
		return root
	}

	return helper(0, 0, len(preorder)-1)
}
