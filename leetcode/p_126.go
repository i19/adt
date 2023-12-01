package leetcode

//126. 单词接龙II
// 按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，
// 一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
// 1. 每对相邻的单词之间仅有单个字母不同。
// 2. 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
// 3. sk == endWord
// 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回 所有 从 beginWord 到 endWord 的 最短转换序列 ，
// 如果不存在这样的转换序列，返回一个空列表。
// 每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。

// 这是回溯（深度优先），可行，但是慢。。。
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string

	var found bool
	for _, v := range wordList {
		if v == endWord {
			found = true
			break
		}
	}
	if !found {
		return res
	}

	selected := make([]bool, len(wordList))
	var helper func(path []string)
	diffMoreThanOne := func(a, b string) bool {
		diffC := 0
		for j := 0; j < len(a); j++ {
			if a[j] != b[j] {
				diffC++
				if diffC > 1 {
					return true
				}
			}
		}
		return false
	}
	helper = func(path []string) {
		if len(res) > 0 && len(path) > len(res[0]) {
			return
		}

		lastWord := path[len(path)-1]
		if lastWord == endWord {
			if len(res) == 0 || len(res[0]) == len(path) {
				res = append(res, append([]string{}, path...))
			} else if len(res[0]) > len(path) {
				res = append([][]string{}, append([]string{}, path...))
			}
			return
		}

		for i := 0; i < len(wordList); i++ {
			if selected[i] {
				continue
			}
			if diffMoreThanOne(lastWord, wordList[i]) {
				continue
			}

			selected[i] = true
			helper(append(path, wordList[i]))
			selected[i] = false
		}
	}
	helper([]string{beginWord})
	return res
}

// 如果是为了寻找 最短，那么应该是广度优先，由于需要返回路径，所以增加一个回溯将 path 进行保存
// func findLaddersII(beginWord string, endWord string, wordList []string) [][]string {
// 	var res [][]string
// 	var found bool
// 	for _, v := range wordList {
// 		if v == endWord {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		return res
// 	}

// 	diffMoreThanOne := func(a, b string) bool {
// 		diffC := 0
// 		for j := 0; j < len(a); j++ {
// 			if a[j] != b[j] {
// 				diffC++
// 				if diffC > 1 {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}

// }
