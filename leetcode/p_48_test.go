package leetcode

import (
	"fmt"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	// fmt.Println(rotate(3, 0, 0))
	// fmt.Println(rotate(3, 0, 2))
	// fmt.Println(rotate(3, 2, 2))
	// fmt.Println(rotate(3, 2, 0))
	// fmt.Println(rotate(3, 0, 1))

	rotate(6)
	fmt.Println(rotate2(6, 2, 2))
}
