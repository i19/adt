package own

import (
	"fmt"
	"testing"
)

func Test_preOrderII(t *testing.T) {
	tree := &treeNode{
		Val: 1,
		Left: &treeNode{
			Val: 7,
			Left: &treeNode{
				Val: 4,
			},
			Right: &treeNode{
				Val: 5,
			},
		},
		Right: &treeNode{
			Val: 3,
			Left: &treeNode{
				Val: 6,
			},
			Right: &treeNode{
				Val: 7,
			},
		},
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	preOrderII(tree, &res, &path)
	fmt.Println("path from root to 7", res)

	res = make([][]int, 0)
	path = make([]int, 0)
	preOrderIII(tree, &res, &path)
	fmt.Println("path from root to 7 bypass 3", res)
}

func Test_backtrace(t *testing.T) {
	tree := &treeNode{
		Val: 1,
		Left: &treeNode{
			Val: 7,
			Left: &treeNode{
				Val: 4,
			},
			Right: &treeNode{
				Val: 5,
			},
		},
		Right: &treeNode{
			Val: 3,
			Left: &treeNode{
				Val: 6,
			},
			Right: &treeNode{
				Val: 7,
			},
		},
	}

	state := make([]*treeNode, 0)
	choice := []*treeNode{tree}
	res := make([][]int, 0)
	backtrackIII(&state, &choice, &res)
	fmt.Println("path from root to 7 bypass 3", res)
}

func Test_permutationsI(t *testing.T) {
	// 无重复元素
	// nums := []int{1, 2, 5}
	// res := permutationsI(nums)
	// fmt.Println(res)

	// res = permutationsII(nums)
	// fmt.Println(res)
	// 有重复元素
	nums := []int{1, 1, 1, 5, 5}
	fmt.Println(permutationsIII(nums))
}
