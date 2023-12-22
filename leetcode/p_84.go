package leetcode

// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram
func largestRectangleArea(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 {
			if heights[i] < heights[stack[len(stack)-1]] {
				right[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) != 0 {
			if heights[i] < heights[stack[len(stack)-1]] {
				left[stack[len(stack)-1]] = stack[len(stack)-1] - i
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < len(heights); i++ {
		width := 0
		if left[i] == 0 {
			width = i + 1
		} else {
			width = left[i]
		}

		if right[i] == 0 {
			width += (len(heights) - 1 - i)
		} else {
			width = width + right[i] - 1
		}

		res = max(
			res,
			width*heights[i],
		)
	}
	return res
}

func largestRectangleAreaII(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		left[i] = i
		right[i] = len(heights) - 1 - i
	}

	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
			right[stack[len(stack)-1]] = i - stack[len(stack)-1] - 1
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
			left[stack[len(stack)-1]] = stack[len(stack)-1] - i - 1
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	res := 0
	for i := 0; i < len(heights); i++ {
		res = max(res, (right[i]+left[i]+1)*heights[i])
	}
	return res
}
