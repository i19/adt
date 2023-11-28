package leetcode

// 79. 单词搜索
// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {
	state := [][]int{}
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	var helper func() bool
	helper = func() bool {
		if len(state) == len(word) {
			return true
		}
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[0]); col++ {
				// 同一个元素不可使用两次
				if used[row][col] {
					continue
				}
				// 字符需要相等
				if word[len(state)] != board[row][col] {
					continue
				}
				// 需要判断相邻
				if len(state) > 0 && abs(state[len(state)-1][0]-row)+abs(state[len(state)-1][1]-col) != 1 {
					continue
				}

				state = append(state, []int{row, col})
				used[row][col] = true
				if helper() {
					return true
				} else {
					state = state[:len(state)-1]
					used[row][col] = false
				}
			}
		}
		return false
	}
	return helper()
}

// 这个是第一版本优化，但是好像效果不好
func existII(board [][]byte, word string) bool {
	state := make([][]int, 0)
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	var helper func(int, int) bool
	helper = func(rowS, colS int) bool {
		if len(state) == len(word) {
			return true
		}
		for _, p := range [][]int{
			{rowS - 1, colS},
			{rowS, colS - 1},
			{rowS, colS + 1},
			{rowS + 1, colS},
		} {
			row, col := p[0], p[1]
			// 坐标合法
			if !(row >= 0 && row < len(board) && col >= 0 && col < len(board[0])) {
				continue
			}
			// 同一个元素不可使用两次
			if used[row][col] {
				continue
			}
			// 字符需要相等
			if word[len(state)] != board[row][col] {
				continue
			}

			state = append(state, []int{row, col})
			used[row][col] = true
			if helper(row, col) {
				return true
			} else {
				state = state[:len(state)-1]
				used[row][col] = false
			}
		}

		return false
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == word[0] {
				used[row][col] = true
				state = append(state, []int{row, col})
				if helper(row, col) {
					return true
				}
				used[row][col] = false
				state = make([][]int, 0)
			}
		}
	}

	return helper(0, 0)
}

// 第二版本优化()
// 相对于第一版
// 1 没有保留中间路径具体结果
// 2 将比对逻辑放在了 herlper 内部，使 main 函数更干净一些（
// ！！！但是也要主要 helper 内，将这个检查放在了最前面，是做的状态检查，而不是剪枝 选择

func existIII(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	var helper func(rowS, colS, pos int) bool
	helper = func(rowS, colS, pos int) bool {
		// 路径错误
		if board[rowS][colS] != word[pos] {
			return false
		}

		// 找到结果
		if pos == len(word)-1 {
			return true
		}
		// 尝试解题
		for _, p := range [][]int{
			{rowS, colS - 1},
			{rowS, colS + 1},
			{rowS - 1, colS},
			{rowS + 1, colS},
		} {
			row, col := p[0], p[1]
			// 坐标合法
			if !(row >= 0 && row < len(board) && col >= 0 && col < len(board[0])) {
				continue
			}
			// 同一个元素不可使用两次
			if used[row][col] {
				continue
			}
			used[row][col] = true
			if helper(row, col, pos+1) {
				return true
			} else {
				used[row][col] = false
			}
		}

		return false
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			used[row][col] = true
			if helper(row, col, 0) {
				return true
			}
			used[row][col] = false
		}
	}
	return false
}
