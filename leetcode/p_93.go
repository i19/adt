package leetcode

import (
	"strconv"
	"strings"
)

// 93. 复原 IP 地址
// https://leetcode.cn/problems/restore-ip-addresses/
func restoreIpAddresses(s string) (res []string) {
	tres := make([]string, 4)
	var dfs func(start, pos int)
	dfs = func(start, pos int) {
		if pos == 4 {
			if start == len(s) {
				res = append(res, strings.Join(tres, "."))
			}
			return
		}
		for i := 1; i <= 3 && start+i <= len(s); i++ {
			ipS := s[start : start+i]
			if ipS[0] == '0' && len(ipS) > 1 {
				return
			}
			ipN, _ := strconv.Atoi(ipS)
			if ipN >= 0 && ipN <= 255 {
				tres[pos] = ipS
				dfs(start+i, pos+1)
			}
		}
	}
	dfs(0, 0)
	return
}
