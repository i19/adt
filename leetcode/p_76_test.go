package leetcode

import (
	"fmt"
	"testing"
)

func Test_minWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	x := "ABC"
	// s = "AB"
	// x = "B"
	fmt.Println(minWindow(s, x))
}
