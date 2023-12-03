package leetcode

// 73. 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes
// 空间复杂度 O(m+n)
func setZeroes(matrix [][]int) {
	rm := make(map[int]bool)
	cm := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rm[i] = true
				cm[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		if rm[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
			continue
		}
		for k := range cm {
			matrix[i][k] = 0
		}
	}
}
