package leetcode

// 994. 腐烂的橘子
// https://leetcode.cn/problems/rotting-oranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var start [][]int
	good := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				start = append(start, []int{i, j})
			} else if grid[i][j] == 1 {
				good++
			}
		}
	}
	if good == 0 {
		return 0
	}

	t := 0
	for {
		start = marker(start, grid)
		good -= len(start)
		if len(start) != 0 {
			t++
		} else {
			break
		}
	}

	if good == 0 {
		return t
	} else {
		return -1
	}
}

func marker(start [][]int, grid [][]int) (res [][]int) {
	rows, cols := len(grid), len(grid[0])
	for _, pos := range start {
		r, c := pos[0], pos[1]
		tbds := [][]int{
			{r - 1, c},
			{r + 1, c},
			{r, c - 1},
			{r, c + 1},
		}

		for _, tbd := range tbds {
			tbdR, tbdC := tbd[0], tbd[1]
			if tbdR >= 0 && tbdR < rows && tbdC >= 0 && tbdC < cols && grid[tbdR][tbdC] == 1 {
				grid[tbdR][tbdC] = 2
				res = append(res, []int{tbdR, tbdC})
			}
		}
	}
	return
}
