package own

import "fmt"

func printCube(points [][]int, matrix int) {
	mP := make(map[string]struct{})
	for _, p := range points {
		mP[fmt.Sprintf("%d%d", p[0], p[1])] = struct{}{}
	}
	for i := 0; i < matrix; i++ {
		for j := 0; j < matrix; j++ {
			fmt.Print("[")
			if _, ok := mP[fmt.Sprintf("%d%d", i, j)]; ok {
				fmt.Print("Q")
			} else {
				fmt.Print("#")
			}
			fmt.Print("]")
		}
		println()
	}
}

// N 皇后问题
// 根据国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
// 给定 n 个皇后和一个  n x n 大小的棋盘，寻找使得所有皇后之间无法相互攻击的摆放方案。
func nQueueIsSolution(state *[][]int, n int) bool {
	return len(*state) == n
}

// 除枝
func nQueenIsValid(states *[][]int, choice []int) bool {
	if len(*states) == 0 {
		return true
	}

	for _, state := range *states {
		// 去除同一行 or 同一列
		if choice[0] == state[0] || choice[1] == state[1] {
			return false
		}
		// 去除对角线
		// 将棋盘想象为坐标系，对角线则满足
		// y = x + b 或者 y = -x + b ==>
		// y - x = b 或者 y + x = b
		// 也就是在同一条对角线上的点，会满足 相加或者相减得到的值相等
		if choice[0]-choice[1] == state[0]-state[1] || choice[0]+choice[1] == state[0]+state[1] {
			return false
		}
	}

	return true
}

// 引入 start 限制搜索边界
// 1. 避免相同组合的出现
// 2. 使用 i+1 可以少一次对本行的二次检索
func backTrackNQueenHelper(start int, state *[][]int, res *[][][]int, n int) {
	if nQueueIsSolution(state, n) {
		*res = append(*res, append([][]int{}, *state...))
	}

	for i := start; i < n; i++ {
		for j := 0; j < n; j++ {
			choice := []int{i, j}
			if nQueenIsValid(state, choice) {
				*state = append(*state, choice)
				backTrackNQueenHelper(i+1, state, res, n)
				*state = (*state)[:len(*state)-1]
			}
		}
	}
}

func backTrackNQueen(n int) [][][]int {
	var res [][][]int
	var state [][]int

	backTrackNQueenHelper(0, &state, &res, n)
	return res
}
