package leetcode

// 37. 解数独
// https://leetcode.cn/problems/sudoku-solver

func getPos(row, col int) int {
	return (row/3)*3 + (col / 3)
}

func solveSudoku(board [][]byte) {
	rowFilter := make(map[int]map[int]bool)
	colFilter := make(map[int]map[int]bool)
	cubFilter := make(map[int]map[int]bool)
	remain := 0
	for i := 0; i < 9; i++ {
		rowFilter[i] = make(map[int]bool, 0)
		colFilter[i] = make(map[int]bool, 0)
		cubFilter[i] = make(map[int]bool, 0)
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			v := board[row][col]
			if v != '.' {
				rowFilter[row][int(v-'0')] = true
				colFilter[col][int(v-'0')] = true
				// cubFilter[row/3+col/3][int(v-'0')] = true
				cubFilter[getPos(row, col)][int(v-'0')] = true
			} else {
				remain++
			}
		}
	}
	var helper func(board [][]byte, rowFilter, colFilter, cubFilter map[int]map[int]bool, remain int) bool
	helper = func(board [][]byte, rowFilter, colFilter, cubFilter map[int]map[int]bool, remain int) bool {
		if remain == 0 {
			return true
		}

		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for choice := 1; choice <= 9; choice++ {
						// 剪枝
						if !(rowFilter[row][choice] || colFilter[col][choice] || cubFilter[getPos(row, col)][choice]) {
							// 做选择
							board[row][col] = byte('0' + choice)
							rowFilter[row][choice] = true
							colFilter[col][choice] = true
							cubFilter[getPos(row, col)][choice] = true
							remain--
							// 如果解决失败则回溯，否则返回成功
							if !helper(board, rowFilter, colFilter, cubFilter, remain) {
								// 回溯
								board[row][col] = '.'
								remain++
								rowFilter[row][choice] = false
								colFilter[col][choice] = false
								cubFilter[getPos(row, col)][choice] = false
							} else {
								return true
							}
						}
					}
					return false
				}
			}
		}
		return false
	}
	helper(board, rowFilter, colFilter, cubFilter, remain)
}

func solveSudokuII(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

// 这是借鉴 leetcode 解法调整的
// 相对于 solveSudoku 中使用字典，这里调整成为数组，速度会更快一些
func solveSudokuIII(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}
	var helper func(pos int) bool
	helper = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}

		row, col := spaces[pos][0], spaces[pos][1]
		for choice := 0; choice < 9; choice++ {
			if !(line[row][choice] || column[col][choice] || block[row/3][col/3][choice]) {
				board[row][col] = byte('1' + choice)
				line[row][choice] = true
				column[col][choice] = true
				block[row/3][col/3][choice] = true

				if helper(pos + 1) {
					return true
				}
				board[row][col] = '.'
				line[row][choice] = false
				column[col][choice] = false
				block[row/3][col/3][choice] = false
			}
		}

		return false
	}
	helper(0)
}
