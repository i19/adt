package leetcode

// 85. 最大矩形
// https://leetcode.cn/problems/maximal-rectangle
// 借鉴 84 题中柱形图求最大面积
func helper84(heights []int) (area int) {
	col := len(heights)
	left, right := make([]int, col), make([]int, col)
	for i := 0; i < col; i++ {
		left[i] = -1
		right[i] = col
	}
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for i := col - 1; i >= 0; i-- {
		for len(stack1) > 0 && heights[i] < heights[stack1[len(stack1)-1]] {
			left[stack1[len(stack1)-1]] = i
			stack1 = stack1[:len(stack1)-1]
		}
		stack1 = append(stack1, i)
	}
	for i := 0; i < col; i++ {
		for len(stack2) > 0 && heights[i] < heights[stack2[len(stack2)-1]] {
			right[stack2[len(stack2)-1]] = i
			stack2 = stack2[:len(stack2)-1]
		}
		stack2 = append(stack2, i)
	}
	for i := 0; i < col; i++ {
		area = max(area, heights[i]*(right[i]-left[i]-1))
	}
	return
}

func maximalRectangleI(matrix [][]byte) (ans int) {
	row, col := len(matrix), len(matrix[0])
	_matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		_matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				_matrix[i][j] = 1
				ans = 1
			}
		}
	}
	if ans == 0 {
		return
	}
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if _matrix[i][j] == 1 {
				_matrix[i][j] += _matrix[i-1][j]
			}
		}
	}

	for _, heights := range _matrix {
		ans = max(ans, helper84(heights))
	}
	return
}

func maximalRectangleII(matrix [][]byte) (ans int) {
	row, col := len(matrix), len(matrix[0])
	heights := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, helper84(heights))
	}
	return
}
