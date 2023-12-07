package leetcode

// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands
func numIslands(grid [][]byte) int {
	marked := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		marked[i] = make([]bool, len(grid[0]), len(grid[0]))
	}
	mark := func(r, c int) {
		tbd := [][]int{{r, c}}
		for {
			l := len(tbd)
			if l == 0 {
				return
			}
			for _, p := range tbd {
				x, y := p[0], p[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
					if !marked[x][y] && grid[x][y] == '1' {
						marked[x][y] = true
						tbd = append(tbd, [][]int{
							{x - 1, y},
							{x + 1, y},
							{x, y - 1},
							{x, y + 1},
						}...)
					}
				}
			}
			tbd = tbd[l:]
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !marked[i][j] {
				mark(i, j)
				res++
			}
		}
	}
	return res
}

// 更简洁的写法
func numIslandsII(grid [][]byte) int {
	var ans int
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans += 1
			}
		}
	}
	return ans
}
