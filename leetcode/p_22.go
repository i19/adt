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
