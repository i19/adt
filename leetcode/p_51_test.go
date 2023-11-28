package leetcode

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	n := 20
	// results := solveNQueensII(n)
	results := solveNQueensII(n)
	fmt.Println(len(results))
}
