package leetcode

import (
	"slices"
)

// 双扩散
func pathWithObstacles(obstacleGrid [][]int) (ans [][]int) {
	rn, cn := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[rn-1][cn-1] == 1 {
		return
	}
	if rn == 1 && cn == 1 && obstacleGrid[0][0] == 0 {
		ans = [][]int{{0, 0}}
		return
	}

	q1 := [][]int{{0, 0}}
	q2 := [][]int{{rn - 1, cn - 1}}
	q1Mark, q2Mark := 2, 3
	obstacleGrid[0][0] = q1Mark
	obstacleGrid[rn-1][cn-1] = q2Mark

	from := make([][][]int, rn)
	for i := 0; i < rn; i++ {
		from[i] = make([][]int, cn)
	}

	prev, pivot := []int{}, []int{}
	buildPath := func(p []int) [][]int {
		path := [][]int{}
		fromP := p
		for fromP != nil {
			path = append(path, fromP)
			if fromP[0] == 0 && fromP[1] == 0 {
				slices.Reverse(path)
			}
			fromP = from[fromP[0]][fromP[1]]
		}
		return path
	}

	bfs := func(q *[][]int, mark int, findMark int, actions [][]int) bool {
		newQ := [][]int{}
		for _, lastP := range *q {
			for _, action := range actions {
				newR, newC := lastP[0]+action[0], lastP[1]+action[1]
				if newR >= 0 && newR < rn && newC >= 0 && newC < cn {
					if obstacleGrid[newR][newC] == 0 {
						obstacleGrid[newR][newC] = mark
						from[newR][newC] = lastP
						newQ = append(newQ, []int{newR, newC})
					} else if obstacleGrid[newR][newC] == findMark {
						prev = lastP
						pivot = []int{newR, newC}
						return true
					}
				}
			}
		}
		*q = newQ
		return false
	}

	for {
		if len(q1) == 0 || len(q2) == 0 {
			return
		}

		if bfs(&q1, q1Mark, q2Mark, [][]int{{1, 0}, {0, 1}}) ||
			bfs(&q2, q2Mark, q1Mark, [][]int{{-1, 0}, {0, -1}}) {
			path1 := buildPath(prev)
			path2 := buildPath(pivot)
			if path1[0][0] == 0 && path1[0][1] == 0 {
				ans = append(path1, path2...)
			} else {
				ans = append(path2, path1...)
			}
			return
		}
	}
}

// 动态规划
func pathWithObstaclesI(obstacleGrid [][]int) (ans [][]int) {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]bool, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]bool, col)
	}
	dp[0][0] = obstacleGrid[0][0] == 0
	for i := 1; i < row && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = dp[i-1][0]
	}
	for i := 1; i < col && obstacleGrid[0][i-1] == 0; i++ {
		dp[0][i] = dp[0][i-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col && obstacleGrid[i][j] == 0; j++ {
			dp[i][j] = dp[i][j-1] || dp[i-1][j]
		}
	}

	if !dp[row-1][col-1] {
		return
	}

	ans = make([][]int, row+col-1)
	ans[len(ans)-1] = []int{row - 1, col - 1}
	for pos := len(ans) - 2; pos >= 0; pos-- {
		last := ans[pos+1]
		for _, dir := range [][]int{{-1, 0}, {0, -1}} {
			newR, newC := last[0]+dir[0], last[1]+dir[1]
			if newR >= 0 && newC >= 0 && dp[newR][newC] {
				ans[pos] = []int{newR, newC}
				break
			}
		}
	}
	return
}
