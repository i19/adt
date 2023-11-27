package leetcode

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	// candidates := []int{10, 1, 2, 7, 6, 1, 5}
	candidates := []int{1, 1, 1, 5, 2, 6, 2}
	target := 8
	for _, r := range combinationSum2(candidates, target) {
		fmt.Println(r)
	}
}
