package booking

import (
	"sort"
)

// 2146. 价格范围内最高排名的 K 样物品
func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) (ans [][]int) {
	rowN, colN := len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	visited := make([][]bool, rowN)
	for i := 0; i < rowN; i++ {
		visited[i] = make([]bool, colN)
	}

	q := [][]int{start}
	visited[start[0]][start[1]] = true
	for len(q) != 0 {
		l := len(q)
		sort.Slice(q, func(i, j int) bool {
			ir, ic, jr, jc := q[i][0], q[i][1], q[j][0], q[j][1]
			pi, pj := grid[ir][ic], grid[jr][jc]
			return (pi < pj) || (pi == pj && ir < jr) || (pi == pj && ir == jr && ic < jc)
		})
		left := sort.Search(l, func(i int) bool { return grid[q[i][0]][q[i][1]] >= low })
		right := sort.Search(l, func(i int) bool { return grid[q[i][0]][q[i][1]] > high })
		ans = append(ans, q[left:right]...)
		if len(ans) >= k {
			ans = ans[:k]
			return
		}

		for _, p := range q {
			for _, dir := range directions {
				nr, nc := p[0]+dir[0], p[1]+dir[1]
				if nr >= 0 && nr < rowN && nc >= 0 && nc < colN && grid[nr][nc] != 0 && !visited[nr][nc] {
					q = append(q, []int{nr, nc})
					visited[nr][nc] = true
				}
			}
		}
		q = q[l:]
	}

	return
}
