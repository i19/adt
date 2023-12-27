package leetcode

import (
	"strings"
)

// 51. N 皇后
// https://leetcode.cn/problems/n-queens/
func solveNQueensHelperFill(size, pos int) string {
	r := make([]rune, size)
	for i := 0; i < size; i++ {
		if i == pos {
			r[i] = 'Q'
		} else {
			r[i] = '.'
		}
	}
	return string(r)
}

func solveNQueensI(n int) [][]string {
	state := make([]int, 0)
	res := make([][]string, 0)
	colFilter := make(map[int]bool, n)
	positiveFilter := make(map[int]bool, n*2)
	negativeFilter := make(map[int]bool, n*2)
	solveNQueensIHelper(&state, &res, n, 0, colFilter, positiveFilter, negativeFilter)
	return res
}
func solveNQueensIHelper(state *[]int, res *[][]string, n, startRow int,
	colFilter map[int]bool, positiveFilter map[int]bool, negativeFilter map[int]bool) {
	if len(*state) == n {
		r := make([]string, len(*state))
		for i, v := range *state {
			r[i] = solveNQueensHelperFill(len(*state), v)
		}
		*res = append(*res, r)
	}

	for row := startRow; row < n; row++ {
		for col := 0; col < n; col++ {
			if !colFilter[col] && !positiveFilter[row+col] && !negativeFilter[row-col] {
				*state = append(*state, col)
				colFilter[col] = true
				positiveFilter[row+col] = true
				negativeFilter[row-col] = true
				solveNQueensIHelper(state, res, n, row+1, colFilter, positiveFilter, negativeFilter)
				*state = (*state)[:len(*state)-1]
				colFilter[col] = false
				positiveFilter[row+col] = false
				negativeFilter[row-col] = false
			}
		}
	}
}

// 相对于 solveNQueensI 将 filter 从 map 修改成 数组
// 注意，在 helper 中 nagativeFilter 的负数问题，通过 增加 n 进行了规避
func solveNQueensII(n int) [][]string {
	state := make([]int, 0)
	res := make([][]string, 0)
	colFilter := make([]bool, n)
	positiveFilter := make([]bool, n*2)
	negativeFilter := make([]bool, n*2)
	solveNQueensIIHelper(&state, &res, n, 0, colFilter, positiveFilter, negativeFilter)
	return res
}
func solveNQueensIIHelper(state *[]int, res *[][]string, n, startRow int,
	colFilter []bool, positiveFilter []bool, negativeFilter []bool) {
	if len(*state) == n {
		r := make([]string, len(*state))
		for i, v := range *state {
			r[i] = solveNQueensHelperFill(len(*state), v)
		}
		*res = append(*res, r)
	}

	for row := startRow; row < n; row++ {
		for col := 0; col < n; col++ {
			if !colFilter[col] && !positiveFilter[row+col] && !negativeFilter[row-col+n] {
				*state = append(*state, col)
				colFilter[col] = true
				positiveFilter[row+col] = true
				negativeFilter[row-col+n] = true
				solveNQueensIIHelper(state, res, n, row+1, colFilter, positiveFilter, negativeFilter)
				*state = (*state)[:len(*state)-1]
				colFilter[col] = false
				positiveFilter[row+col] = false
				negativeFilter[row-col+n] = false
			}
		}
	}
}

func solveNQueensIII(n int) (res [][]string) {
	posOfCol := make([]int, n)

	sumFilter := make([]bool, 2*n+1)
	minFilter := make([]bool, 2*n+1)
	colFilter := make([]bool, n)

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tres := make([]string, n)
			for c, r := range posOfCol {

				tres[c] = strings.Repeat(".", r) + "Q" + strings.Repeat(".", n-1-r)
			}
			res = append(res, tres)
		}
		for col, filtered := range colFilter {
			minusK := col - row + n
			if !filtered && !sumFilter[col+row] && !minFilter[minusK] {
				posOfCol[col] = row
				colFilter[col], sumFilter[col+row], minFilter[minusK] = true, true, true
				dfs(row + 1)
				colFilter[col], sumFilter[col+row], minFilter[minusK] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
