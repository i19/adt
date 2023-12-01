package leetcode

// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	next := make(map[int]int)
	prev := make(map[int]int)
	for _, num := range nums {
		if _, ok := next[num-1]; ok {
			prev[num] = num - 1
			next[num-1] = num
		}
		if _, ok := next[num+1]; ok {
			prev[num+1] = num
			next[num] = num + 1
		}
		if _, ok := next[num]; !ok {
			next[num] = num
		}
		if _, ok := prev[num]; !ok {
			prev[num] = num
		}
	}

	var heads []int
	for k, v := range prev {
		if k == v {
			heads = append(heads, k)
		}
	}

	maxL := 1
	for _, head := range heads {
		tl := 1
		for head != next[head] {
			tl++
			head = next[head]
			if tl > maxL {
				maxL = tl
			}
		}

	}

	return maxL
}

// leetcode 的解法，更优雅
// 从小到大
func longestConsecutiveII(nums []int) int {
	m := make(map[int]bool)
	for _, k := range nums {
		m[k] = true
	}
	maxL := 0
	for k := range m {
		if !m[k-1] {
			cn := k
			tl := 1
			for m[cn+1] {
				cn++
				tl++
			}
			if tl > maxL {
				maxL = tl
			}
		}
	}
	return maxL
}

// 从大到小
func longestConsecutiveIIII(nums []int) int {
	m := make(map[int]bool)
	for _, k := range nums {
		m[k] = true
	}
	maxL := 0
	for k := range m {
		if !m[k+1] {
			cn := k
			tl := 1
			for m[cn-1] {
				cn--
				tl++
			}
			if tl > maxL {
				maxL = tl
			}
		}
	}
	return maxL
}
