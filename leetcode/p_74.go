package leetcode

// 74. 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix
func searchMatrix(matrix [][]int, target int) bool {
	rowN, colN := len(matrix), len(matrix[0])
	total := rowN * colN

	left, right := 0, total-1
	for left <= right {
		mid := left + (right-left)/2
		mv := matrix[mid/colN][mid%colN]
		if mv == target {
			return true
		} else if mv > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
