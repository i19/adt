package own

import (
	"container/list"
	"fmt"
)

// 广度优先
func levelOrder(root *treeNode) {
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		node := q.Remove(q.Front()).(*treeNode)
		fmt.Printf("%d ", node.Val)

		if node.Left != nil {
			q.PushBack(node.Left)
		}

		if node.Right != nil {
			q.PushBack(node.Right)
		}
	}
}

// 前序
// 根节点 -> 左节点 -> 右节点
func frontOrder(node *treeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Val)
	frontOrder(node.Left)
	frontOrder(node.Right)
}

// 中序
// 左节点 -> 根节点 -> 右节点
func middleOrder(node *treeNode) {
	if node == nil {
		return
	}

	middleOrder(node.Left)
	fmt.Printf("%d ", node.Val)
	middleOrder(node.Right)
}

// 后序
// 左节点 -> 右节点 -> 根节点
func afterOrder(node *treeNode) {
	if node == nil {
		return
	}

	afterOrder(node.Left)
	afterOrder(node.Right)
	fmt.Printf("%d ", node.Val)
}

// binary tree ops
func (node *treeNode) search(val int) *treeNode {
	x := node
	for x != nil {
		if x.Val == val {
			break
		} else if x.Val > val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

func (node *treeNode) insert(val int) *treeNode {
	if node == nil {
		return &treeNode{Val: val}
	}

	pre := node
	for curr := node; curr != nil; {
		pre = curr
		if curr.Val == val {
			return node
		} else if curr.Val < val {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	if pre.Val > val {
		pre.Left = &treeNode{
			Val: val,
		}
	} else {
		pre.Right = &treeNode{
			Val: val,
		}
	}

	return node

}
