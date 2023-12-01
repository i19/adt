package leetcode

// 127. 单词接龙
// https://leetcode.cn/problems/word-ladder/
// 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
// 每一对相邻的单词只差一个字母。
// 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
// sk == endWord
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，
// 返回从 beginWord 到 endWord 的 最短转换序列的单词数目。如果不存在这样的转换序列，返回 0 。

// 从端点较少侧进行广度优先，是不是会比较快
func diffMoreThantOne(a, b string) bool {
	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c++
		}
		if c > 1 {
			return true
		}
	}
	return false
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	type layer struct {
		mark     int
		distance int
		cp, np   []string
	}

	found := false
	mark := make([]int, len(wordList))
	for i, v := range wordList {
		if v == endWord {
			mark[i] = -1
			found = true
			break
		}
		if v == beginWord {
			mark[i] = 1
		}
	}
	if !found {
		return 0
	}

	fromB := &layer{distance: 1, mark: 1, cp: make([]string, 1, len(wordList)), np: make([]string, 0, len(wordList))}
	fromD := &layer{distance: 1, mark: -1, cp: make([]string, 1, len(wordList)), np: make([]string, 0, len(wordList))}
	fromB.cp[0] = beginWord
	fromD.cp[0] = endWord
	x := fromB
	for len(x.cp) != 0 {
		if len(fromB.cp) > len(fromD.cp) {
			x = fromD
		}

		for i := 0; i < len(wordList); i++ {
			if mark[i] != x.mark {
				for _, p := range x.cp {
					if !diffMoreThantOne(wordList[i], p) {
						if mark[i] == 0 {
							// 该节点未被扫描过
							mark[i] = x.mark
							x.np = append(x.np, wordList[i])
							// 标记以后，就不用在计算了
							break
						} else {
							//该节点被对向标记了
							return fromB.distance + fromD.distance
						}
					}
				}
			}
		}
		x.distance++
		x.cp, x.np = x.np, x.cp[:0]
	}

	return 0
}
