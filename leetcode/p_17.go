package leetcode

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number
func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return
	}

	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	comb := make([]rune, len(digits))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(comb))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			comb[i] = c
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
