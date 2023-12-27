package leetcode

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	n := 4
	// results := solveNQueensII(n)
	results := solveNQueensIII(n)
	fmt.Println(len(results))
}
