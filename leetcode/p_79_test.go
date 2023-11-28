package leetcode

import "testing"

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	// board = [][]byte{
	// 	{'a', 'a', 'b', 'a', 'a', 'b'},
	// 	{'a', 'a', 'b', 'b', 'b', 'a'},
	// 	{'a', 'a', 'a', 'a', 'b', 'a'},
	// 	{'b', 'a', 'b', 'b', 'a', 'b'},
	// 	{'a', 'b', 'b', 'a', 'b', 'a'},
	// 	{'b', 'a', 'a', 'a', 'a', 'b'},
	// }

	// word = "bbbaabbbbbab"

	// board = [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word = "ABCB"

	// board = [][]byte{{'a'}}
	// word = "a"

	println(existIII(board, word))
	println(existII(board, word))
}
