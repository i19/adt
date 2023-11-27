package leetcode

// 外观数列
// https://leetcode.cn/problems/count-and-say
func countAndSay(n int) string {
	r := []rune{'1'}
	for i := 1; i < n; i++ {
		r2 := []rune{'0', r[0]}
		for _, x := range r {
			if x == r2[len(r2)-1] {
				r2[(len(r2)-2)]++
			} else {
				r2 = append(r2, []rune{'1', x}...)
			}
		}
		r = r2
	}
	return string(r)
}
