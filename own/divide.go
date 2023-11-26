package own

// i 代表该树的 根节点 在 preorder 中的位置
// l，r 代表该树在 inorder 中的左右边界
// 这个算法的思路就是，利用 inorder 计算子树节点数量，在配合根节点得出两个子树的边界，每次都求出子树的 根节点
// 然后通过 归 的方式，将树进行组装
func buildTreeHelper(preOrder []int, inorderMap map[int]int, i, l, r int) *treeNode {
	// 说明没有了，要为空了
	if r-l < 0 {
		return nil
	}

	// 该树的根节点位置
	root := &treeNode{Val: preOrder[i]}
	// 该树的根节点在中序排列中的索引
	m := inorderMap[preOrder[i]]
	// i + 1 为 左子树的根节点
	// l 为左子树在 中序排列 中的左边界
	// m - 1 为 左子树在 中序排列 中的右边界
	root.Left = buildTreeHelper(preOrder, inorderMap, i+1, l, m-1)
	// i + 1 + m - l
	// m 与 l 都是 inorder 中的下标，m - l 为左子树的节点数量，所以 i + 1 + m - l 就是右子树的根节点
	// m + 1 为右子树在 中序排列 中的左边界
	// r 为右子树在 中序排列 中的右边界
	root.Right = buildTreeHelper(preOrder, inorderMap, i+1+m-l, m+1, r)
	return root
}

func buildTree(preOrder, inOrder []int) *treeNode {
	// 前序遍历可以确定根节点
	// 中序遍历，可以确认左子树和右子树的元素
	inOrderMap := make(map[int]int)
	for i, v := range inOrder {
		inOrderMap[v] = i
	}

	root := buildTreeHelper(preOrder, inOrderMap, 0, 0, len(inOrder)-1)
	return root
}
