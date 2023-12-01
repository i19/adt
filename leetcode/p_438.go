package leetcode

import (
	"sort"
)

// 38. 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string
// 暴力，如果字符较长，时间会很高
func findAnagramsI(s string, p string) []int {
	pb := []byte(p)
	sort.Slice(pb, func(i, j int) bool { return pb[i] > pb[j] })
	pss := string(pb)

	pl := len(p)
	var r []int
	for i := 0; i <= len(s)-pl; i++ {
		s1 := []byte(s[i : i+pl])
		println(string(s1))
		sort.Slice(s1, func(i, j int) bool { return s1[i] > s1[j] })
		if string(s1) == pss {
			r = append(r, i)
		}
	}
	return r
}

// 滑动窗口，统计奇数法
func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
