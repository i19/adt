package leetcode

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses
func isValid(s string) bool {
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	remain := make([]rune, 0)
	for _, x := range s {
		if _, ok := pair[x]; !ok {
			remain = append(remain, x)
		} else {
			if len(remain) > 0 && remain[len(remain)-1] == pair[x] {
				remain = remain[:len(remain)-1]
			} else {
				return false
			}
		}
	}
	return len(remain) == 0
}
