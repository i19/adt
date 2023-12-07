package leetcode

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var res []int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)

	newHead := &TreeNode{}
	dummy := newHead
	for _, v := range res {
		dummy.Right = &TreeNode{
			Val: v,
		}
		dummy = dummy.Right
	}
	root.Left = nil
	root.Right = newHead.Right.Right
}

func flattenII(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
