package leetcode

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n := 10
	for _, s := range generateParenthesis(n) {
		fmt.Println("可用括号:", s)
	}
}
