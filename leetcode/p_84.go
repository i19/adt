package leetcode

import (
	"math"
)

// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram

// 按宽度枚举
func largestRectangleAreaI(heights []int) (ans int) {
	for left := 0; left < len(heights); left++ {
		minH := math.MaxInt
		for right := left; right < len(heights); right++ {
			minH = min(minH, heights[right])
			ans = max(ans, minH*(right-left+1))
			if ans == 24 {
				println("he llo world")
			}
		}
	}
	return
}

// 按高度枚举（扩散）
// O(n^2),内侧的两个循环加起来为 n
func largestRectangleAreaII(heights []int) (ans int) {
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left > -1 && heights[left] >= heights[i] {
			left--
		}
		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}
		ans = max(ans, heights[i]*(right-left-1))
	}
	return
}

// 双单调栈 O(n)
// 枚举 高度的思路，本质是寻找两侧第一个比自己小的数，那么这个逻辑就适合单调栈实现
func largestRectangleAreaIII(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < len(heights); i++ {
		left[i] = -1
		right[i] = len(heights)
	}
	stack1, stack2 := []int{}, []int{}
	// 左侧第一个比本身小的 O(n)
	for i := n - 1; i >= 0; i-- {
		for len(stack1) != 0 && heights[stack1[len(stack1)-1]] > heights[i] {
			left[stack1[len(stack1)-1]] = i
			stack1 = stack1[:len(stack1)-1]
		}
		stack1 = append(stack1, i)
	}
	// 右侧第一个比本身小的 O(n)
	for i := 0; i < len(heights); i++ {
		for len(stack2) != 0 && heights[stack2[len(stack2)-1]] > heights[i] {
			right[stack2[len(stack2)-1]] = i
			stack2 = stack2[:len(stack2)-1]
		}
		stack2 = append(stack2, i)
	}

	ans := 0
	// 遍历高度，利用单调栈完成发散 O(n)
	for i := 0; i < len(heights); i++ {
		ans = max(ans, heights[i]*(right[i]-left[i]-1))
	}
	return ans
}
