package leetcode

// 54. 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	mark := make([][]bool, m)
	for i := 0; i < len(matrix); i++ {
		mark[i] = make([]bool, n)
	}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	d := 0
	ri, ci := 0, 0
	res := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		res[i] = matrix[ri][ci]
		mark[ri][ci] = true

		nextRi, nextCi := ri+directions[d][0], ci+directions[d][1]
		if nextRi >= 0 && nextRi < m && nextCi >= 0 && nextCi < n && !mark[nextRi][nextCi] {
			ri, ci = nextRi, nextCi
			continue
		} else {
			d = (d + 1) % 4
			ri, ci = ri+directions[d][0], ci+directions[d][1]
		}
	}
	return res
}
