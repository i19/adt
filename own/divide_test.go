package own

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	pre := []int{3, 9, 6, 2, 1, 7}
	ino := []int{6, 9, 3, 1, 2, 7}
	tree := buildTree(pre, ino)
	levelOrder(tree)
	println()
	afterOrder(tree)
}
