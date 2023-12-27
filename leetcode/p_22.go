package leetcode

import "fmt"

// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses
func generateParenthesisHelper(cmb string, size int, res *[]string, lr, rr int) {
	if len(cmb) == size*2 {
		*res = append(*res, cmb)
		return
	}
	if lr > 0 {
		generateParenthesisHelper(fmt.Sprintf("%s(", cmb), size, res, lr-1, rr)
	}
	// 只有 rr > lr 的时候，才可以打印 右括号，这样生成的括号才合法
	if rr > lr {
		generateParenthesisHelper(fmt.Sprintf("%s)", cmb), size, res, lr, rr-1)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisHelper("", n, &res, n, n)
	return res
}

func generateParenthesisII(n int) (res []string) {
	tres := make([]rune, n*2)
	tres[0] = '('
	lc, rc := 1, 0
	var dfs func(pos int)
	dfs = func(pos int) {
		if lc > n || rc > n {
			return
		}

		if pos == len(tres) {
			res = append(res, string(tres))
			return
		}
		if lc > rc {
			tres[pos] = '('
			lc++
			dfs(pos + 1)
			lc--
			tres[pos] = ')'
			rc++
			dfs(pos + 1)
			rc--
		}
		if lc == rc {
			tres[pos] = '('
			lc++
			dfs(pos + 1)
			lc--
		}
	}
	dfs(1)
	return
}

func generateParenthesisIII(n int) (res []string) {
	tres := make([]rune, n*2)
	var dfs func(pos, l, r int)
	dfs = func(pos, l, r int) {
		if pos == len(tres) {
			res = append(res, string(tres))
			return
		}
		if l < n {
			tres[pos] = '('
			dfs(pos+1, l+1, r)
		}
		if r < l {
			tres[pos] = ')'
			dfs(pos+1, l, r+1)
		}
	}
	dfs(0, 0, 0)
	return
}
