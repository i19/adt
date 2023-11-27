package leetcode

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7, 8, 9}
	target := 7
	for _, r := range combinationSum(candidates, target) {
		fmt.Println(r)
	}
}
