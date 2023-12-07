package leetcode

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	// [3,5,1 ,6,2, 0,8,null,null,7,4]
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 1}
	n4 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 2}
	n6 := &TreeNode{Val: 0}
	n7 := &TreeNode{Val: 8}
	n8 := &TreeNode{Val: 7}
	n9 := &TreeNode{Val: 4}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n5.Left = n8
	n5.Right = n9

	z := lowestCommonAncestor(n1, n2, n3)

	if z == nil {
		println("damn ....")
	} else {

		println(z.Val)
	}

	// var r []*TreeNode
	// findPathToNode(n1, n3, &r)
	// println(len(r))
}
