package leetcode

// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == q {
		return p
	}
	if root == p || root == q {
		return root
	}

	var helper func(node, target *TreeNode, res *[]*TreeNode) bool
	helper = func(node, target *TreeNode, res *[]*TreeNode) bool {
		if node == nil {
			return false
		}
		*res = append(*res, node)
		if node == target {
			return true
		}
		if helper(node.Left, target, res) || helper(node.Right, target, res) {
			return true
		}
		*res = (*res)[:len(*res)-1]
		return false
	}
	var pathToP []*TreeNode
	var pathToQ []*TreeNode
	helper(root, q, &pathToQ)
	helper(root, p, &pathToP)

	for i := min(len(pathToP), len(pathToQ)) - 1; i >= 0; i-- {
		if pathToP[i] == pathToQ[i] {
			return pathToP[i]
		}
	}
	return nil
}

// 根据定义，若 root 是 p,q 的 最近公共祖先 ，则只可能为以下情况之一：
// 1. p 和 q 在 root 的子树中，且分列 root 的 两侧（即分别在左、右子树中）；
// 2, p 在 q 的子树下，或者 q 在 p 的子树下

// 递归解析：
// 终止条件：
// 		1. 当越过叶节点，则直接返回 nil ；
// 		2. 当 root 等于 p,q ，则直接返回 root ；
// 递推工作：
// 		开启递归左子节点，返回值记为 left ；
// 		开启递归右子节点，返回值记为 right ；
// 返回值： 根据 left 和 right ，可展开为四种情况；
// 		1. 当 left 和 right 同时为空 ：说明 root 的左 / 右子树中都不包含 p,q ，返回 nil ；
// 		2. 当 left 和 right 同时不为空 ：说明 p,q  分列在 root 的 异侧 （分别在 左 / 右子树），因此 root 为最近公共祖先，返回 root ；
// 		3. 当 left 为空 ，right 不为空 ：p,q 都不在 root 的左子树中，直接返回 right 。具体可分为两种情况：
// 			1. p,q 其中一个在 root 的 右子树 中，此时 right 指向 ppp（假设为 ppp ）；
// 			2. p,q 两节点都在 root 的 右子树 中，此时的 right 指向 最近公共祖先节点 ；
// 		4. 当 left 不为空 ， right 为空 ：与情况 3. 同理；

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorII(root.Left, p, q)
	right := lowestCommonAncestorII(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
