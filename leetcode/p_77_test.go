package leetcode

import (
	"fmt"
	"testing"
)

func Test_combine(t *testing.T) {
	n := 1
	k := 1
	for _, v := range combine(n, k) {
		fmt.Println(v)
	}
}
