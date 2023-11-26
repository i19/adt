package own

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	tree := &treeNode{
		Val: 3,
		Left: &treeNode{
			Val: 1,
			Left: &treeNode{
				Val: 0,
			},
			Right: &treeNode{
				Val: 2,
			},
		},
		Right: &treeNode{
			Val: 5,
			Left: &treeNode{
				Val: 4,
			},
			Right: &treeNode{
				Val: 6,
			},
		},
	}
	println("广度优先")
	levelOrder(tree)
	println()

	println("前序遍历")
	frontOrder(tree)
	println()

	println("中序遍历")
	middleOrder(tree)
	println()

	println("后序遍历")
	afterOrder(tree)
	println()
}

func Test_treeNode_insert(t *testing.T) {
	var tree *treeNode
	tree = tree.insert(3)
	tree = tree.insert(1)
	tree = tree.insert(0)
	tree = tree.insert(2)
	tree = tree.insert(5)
	tree = tree.insert(4)
	tree = tree.insert(6)
	tree = tree.insert(6)

	println("中序遍历")
	middleOrder(tree)
	println()

	x := tree.search(5)
	println("中序遍历")
	middleOrder(x)
	println()
}
