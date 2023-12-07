package leetcode

// 240. 搜索二维矩阵 II
// https://leetcode.cn/problems/search-a-2d-matrix-ii
func searchMatrixII(matrix [][]int, target int) bool {
	rowN, colN := len(matrix), len(matrix[0])
	if matrix[0][0] > target || matrix[rowN-1][colN-1] < target {
		return false
	}
	maxCol := colN
	for row := 0; row < rowN; row++ {
		if matrix[row][0] > target {
			return false
		}
		for col := 0; col < maxCol; col++ {
			v := matrix[row][col]
			if v == target {
				return true
			}
			if v > target {
				maxCol = col
				break
			}
		}
	}

	return false
}
