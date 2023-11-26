package own

import (
	"fmt"
	"testing"
)

func TestGraphList_bfs(t *testing.T) {
	g := &GraphList{
		vertex: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		edge: map[int][]int{
			4: {1, 3, 5, 7},
			6: {3, 7},
			5: {2, 4, 8},
			7: {4, 6, 8},
			8: {5, 7},
			0: {1, 3},
			1: {0, 2, 4},
			3: {0, 4, 6},
			2: {1, 5},
		},
	}

	res := g.bfs()
	println("广度优先遍历（BFS）顶点序列为:")
	fmt.Println(res)

	res = g.dfs()
	println("深度优先遍历（DFS）顶点序列为:")
	fmt.Println(res)
}
