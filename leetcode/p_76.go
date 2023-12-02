package leetcode

// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	ls, lt := len(s), len(t)
	mt := make(map[byte]int)
	ms := make(map[byte]int)

	for i := 0; i < lt; i++ {
		mt[t[i]]++
	}
	for i := 0; i < lt; i++ {
		if _, ok := mt[s[i]]; ok {
			ms[s[i]]++
		}
	}

	match := func() bool {
		if len(mt) != len(ms) {
			return false
		}
		for k, v := range mt {
			if ms[k] < v {
				return false
			}
		}
		return true
	}

	if match() {
		return s[:lt]
	}

	start, length := 0, 0
	for left, right := 0, lt; right < ls; right++ {
		if _, ok := mt[s[right]]; !ok {
			continue
		}

		ms[s[right]]++
		for match() && left <= right {
			tl := right - left + 1

			if length == 0 {
				length = tl
			} else if tl < length {
				length = tl
				start = left
			}

			if _, ok := mt[s[left]]; ok {
				ms[s[left]]--
			}
			left++
		}
	}
	return s[start : start+length]
}
