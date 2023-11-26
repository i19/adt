package own

import (
	"fmt"
	"testing"
)

func Test_fractionalKnapsack(t *testing.T) {
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	capacity := 50

	// 贪心算法
	res := fractionalKnapsack(val, wgt, float64(capacity))
	fmt.Println("不超过背包容量的最大物品价值为", res)
	ht := []int{3, 8, 5, 2, 7, 7, 3, 4}
	fmt.Println("最大面积为", maxArea(ht))

	// 表示一个无向图的邻接矩阵
	graph := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	// 调用Prim算法来生成最小生成树
	primMST(graph)
}
