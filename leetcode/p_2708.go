package leetcode

import "sort"

// 2708. 一个小组的最大实力值
// https://leetcode.cn/problems/maximum-strength-of-a-group
func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	pos, neg := []int{}, []int{}
	zero := false
	for _, v := range nums {
		if v > 0 {
			pos = append(pos, v)
		} else if v == 0 {
			zero = true
		} else {
			neg = append(neg, v)
		}
	}
	sort.Ints(neg)

	if len(pos) == 0 && len(neg) < 2 {
		if zero {
			return int64(0)
		} else {
			return int64(neg[0])
		}
	}
	if len(neg)%2 == 1 {
		neg = neg[:len(neg)-1]
	}
	ans := 1
	pos = append(pos, neg...)
	for _, v := range pos {
		ans *= v
	}

	return int64(ans)
}
