package leetcode

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	// n := []int{1, 2, 3, 4, 5, 6, 7}
	n := []int{1, 2}
	rotate(n, 3)
	fmt.Println(n)
}
