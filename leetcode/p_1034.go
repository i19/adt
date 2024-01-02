package leetcode

// 1034. 边界着色
// https://leetcode.cn/problems/coloring-a-border/
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	rowN, colN := len(grid), len(grid[0])
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	selected := make([][]bool, rowN)
	for i := 0; i < rowN; i++ {
		selected[i] = make([]bool, colN)
	}
	expand := func(i, j int) (ans [][]int) {
		for _, dir := range directions {
			r, c := i+dir[0], j+dir[1]
			if r >= 0 && r < rowN && c >= 0 && c < colN && grid[r][c] == grid[i][j] {
				ans = append(ans, []int{r, c})
			}
		}
		return
	}
	q := [][]int{{row, col}}
	selected[row][col] = true
	border := [][]int{}
	for len(q) != 0 {
		l := len(q)
		for _, p := range q {
			s := expand(p[0], p[1])
			if len(s) != 4 {
				border = append(border, p)
			}
			for _, p2 := range s {
				if !selected[p2[0]][p2[1]] {
					q = append(q, p2)
					selected[p2[0]][p2[1]] = true
				}
			}
		}
		q = q[l:]
	}
	for _, p := range border {
		grid[p[0]][p[1]] = color
	}
	return grid
}
