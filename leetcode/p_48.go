package leetcode

// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image

// GOOD
func getNewPos(matrixSize int, row, col int) []int {
	return []int{col, matrixSize - row - 1}
}
func rotate(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size/2; i++ {
		for j := i; j < size-1-i; j++ {
			c1 := []int{i, j}
			c2 := getNewPos(size, c1[0], c1[1])
			c3 := getNewPos(size, c2[0], c2[1])
			c4 := getNewPos(size, c3[0], c3[1])
			matrix[c1[0]][c1[1]], matrix[c2[0]][c2[1]], matrix[c3[0]][c3[1]], matrix[c4[0]][c4[1]] = matrix[c4[0]][c4[1]], matrix[c1[0]][c1[1]], matrix[c2[0]][c2[1]], matrix[c3[0]][c3[1]]
		}
	}
}
