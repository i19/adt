package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPathToNode(node, target *TreeNode, res *[]*TreeNode, path []*TreeNode) bool {
	if node == nil {
		return false
	}
	path = append(path[:0:0], path...) // 复制 path 的内容
	path = append(path, node)
	if node == target {
		*res = append(*res, path...)
		return true
	}
	if findPathToNode(node.Left, target, res, path) || findPathToNode(node.Right, target, res, path) {
		return true
	}
	return false
}

func findPath(root, target *TreeNode) []*TreeNode {
	var res []*TreeNode
	findPathToNode(root, target, &res, []*TreeNode{})
	return res
}

func main() {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 1}
	n1.Left = n2
	n2.Right = n3

	paths := findPath(n1, n3)
	println(len(paths))
}
