package leetcode

import "sort"

// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water

// 按行求解，获取每一行的水
// 复杂度为 m * n
func trapII(height []int) int {
	h := 1
	r := 0
	for {
		block := 0
		left, right := -1, -1
		for i, v := range height {
			if v >= h {
				block++
				if left == -1 {
					left = i
				} else {
					right = i
				}
			}
		}
		if block > 1 {
			r += (right - left - block + 1)
			h++
		} else {
			return r
		}
	}
}

// 得到两个最高，然后向两侧扩散
// 复杂度为 nlogn
func trapI(height []int) int {
	type trapItem struct {
		index  int
		height int
	}
	if len(height) <= 1 {
		return 0
	}
	x := make([]*trapItem, len(height))
	for i, h := range height {
		x[i] = &trapItem{
			index:  i,
			height: h,
		}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].height > x[j].height
	})

	r := 0

	left, right := min(x[0].index, x[1].index), max(x[0].index, x[1].index)
	r = x[1].height * (right - left - 1)
	for i := left + 1; i < right; i++ {
		r -= height[i]
	}
	for _, element := range x[2:] {
		if element.index < left {
			r += ((left - element.index - 1) * element.height)
			for i := element.index + 1; i < left; i++ {
				r -= height[i]
			}
			left = element.index
		} else if element.index > right {
			r += ((element.index - right - 1) * element.height)
			for i := right + 1; i < element.index; i++ {
				r -= height[i]
			}
			right = element.index
		}
	}
	return r
}

// 按列求解
// 使用动态规划维护每一列左侧和右侧最大值
func trapIII(height []int) int {
	dpl := make([]int, len(height))
	dpr := make([]int, len(height))
	dpl[0] = 0
	dpr[len(height)-1] = 0
	for i := 1; i < len(height); i++ {
		dpl[i] = max(dpl[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		dpr[i] = max(dpr[i+1], height[i+1])
	}

	r := 0
	for i := 0; i < len(height)-1; i++ {
		x := min(dpl[i], dpr[i])
		if height[i] < x {
			r += (x - height[i])
		}
	}

	return r
}
