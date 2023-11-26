package own

import (
	"fmt"
	"math"
	"sort"
)

// 分数背包问题
// 给定 n 个物品的val 和 weight, 每个物品只能选择一次，但可以选择物品的一部分，
// 问在不超过背包容量下背包中物品的最大价值。
// 一道简单的贪心问题
func fractionalKnapsack(val, weight []int, rCap float64) float64 {
	items := make([][]float64, len(val))
	for i := 0; i < len(val); i++ {
		items[i] = []float64{float64(val[i]), float64(weight[i])}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0]/items[i][1] > items[j][0]/items[j][1]
	})

	profit := float64(0)

	for i := 0; i < len(items) && rCap > 0; i++ {
		if rCap > items[i][1] {
			profit += items[i][0]
		} else {
			profit += rCap / items[i][1] * items[i][0]
		}
		rCap -= items[i][1]
	}

	return profit
}

// 最大面积问题（最多装水问题）
// 输入一个数组，数组中的每个元素代表一个垂直隔板的高度。数组中的任意两个隔板，以及它们之间的空间可以组成一个容器。
// 容器的容量等于高度和宽度的乘积（即面积），其中高度由较短的隔板决定，宽度是两个隔板的数组索引之差。
// 请在数组中选择两个隔板，使得组成的容器的容量最大，返回最大容量。
func maxArea(ht []int) int {
	r := 0

	for left, right := 0, len(ht)-1; left < right; {
		tr := min(ht[left], ht[right]) * (right - left)
		if tr > r {
			r = tr
		}
		if ht[left] < ht[right] {
			left++
		} else {
			right--
		}
	}
	return r
}

// prim 算法思路(最小生成树)
// 基于的思想是，如果将一棵 spt 中的一条边 切断，则生成的两棵树仍旧是 spt，而切断的边，是将这两棵 spt 合并成新的 spt 的最短路径
// 过程
// Prim算法是一种用于生成无向图的最小生成树（Minimum Spanning Tree）的贪心算法。最小生成树是一个包含了所有顶点但没有环的子图，其边的权重之和最小。Prim算法的逻辑可以概括为以下步骤：
// 1. 初始化：选择一个任意的起始顶点作为最小生成树的根节点，将该顶点加入最小生成树的集合。同时，创建一个用于记录顶点是否已经包含在最小生成树中的布尔数组（例如，isInSPT[]）。
// 2. 更新顶点的关键值：对于每个与最小生成树集合相邻的顶点（即尚未包含在最小生成树中的顶点），计算连接它们的边的权重。如果该权重小于当前已知的权重（或者初始为无穷大），则更新这个顶点的关键值为新的权重。
// 3. 选择最小关键值：从未包含在最小生成树中的顶点中，选择具有最小关键值的顶点（即，选择最小权重边相邻的顶点）。
// 4. 将选择的顶点加入最小生成树：将上一步中选择的顶点加入最小生成树的集合，同时标记它为已包含。
// 重复步骤2至步骤4：重复以上步骤，直到最小生成树包含了所有的顶点（即，直到所有顶点都被标记为已包含）。
// Prim算法的核心思想是从一个初始顶点开始，逐步选择连接到最小权重边的顶点，构建最小生成树。算法根据当前已包含的顶点和它们的权重，贪心地选择下一个要加入的顶点，以确保生成的树的权重最小。算法的最终结果是一个无环的子图，包含了所有顶点，并且其边的权重之和最小。

// INF 用于表示无穷大的常量，表示两顶点之间没有边
const INF = math.MaxInt

func primMST(graph [][]int) {
	vCount := len(graph)

	// isInSpt 创建一个切片用于存储已包含在最小生成树中的顶点，默认为 false
	// pair 切片用于存储顶点到 已知spt 的位置，也就是 spt 中的节点，用于存储 spt 的结构（根顶点无意义，只是为了 startup）
	// cost 切片用于存储 g-spt 连接到 spt 中的最小连接权重
	// 在 prim 算法中，spt 是一直在变化的过程，所以未包含在 spt 中的节点与 spt 的连接诶可能也会变化，意味着 pair 和 cost 中的值也会一直变化
	isInSpt := make([]bool, vCount)
	pair := make([]int, vCount)
	cost := make([]int, vCount)

	// 初始化cost切片，将其所有值设置为无穷大
	for i := 0; i < vCount; i++ {
		cost[i] = INF
	}

	// 从第一个顶点开始构建最小生成树，但是此时该点尚未进入 spt
	cost[0] = 0
	pair[0] = -1

	// spt 的边数量 = len(vertex) -1
	for edgeCount := 0; edgeCount < vCount-1; edgeCount++ {
		// 选择新的划入 spt 的顶点
		newSPTVertex := findNewSPTVertex(isInSpt, vCount, cost)
		isInSpt[newSPTVertex] = true

		// 遍历未在 spt 中并且和 u 节点相连的点，并更新其与 spt 的距离（使用更小的）
		for vertex := 0; vertex < vCount; vertex++ {
			// 已连接 && 该节点在 V-A 中 && 获取该节点到 V 的所有路径，更新权重值
			if !isInSpt[vertex] && graph[vertex][newSPTVertex] != 0 && graph[vertex][newSPTVertex] < cost[vertex] {
				pair[vertex] = newSPTVertex
				cost[vertex] = graph[vertex][newSPTVertex]
			}
		}
	}

	// 打印最小生成树
	fmt.Println("边  权重")
	for i := 1; i < vCount; i++ {
		fmt.Printf("%d - %d    %d\n", pair[i], i, graph[i][pair[i]])
	}
}

// 找到cost切片中的最小值并返回其索引
func findNewSPTVertex(isInSPT []bool, vCount int, cost []int) int {
	minCost := INF
	minV := -1

	for v := 0; v < vCount; v++ {
		if !isInSPT[v] && cost[v] < minCost {
			minV = v
			minCost = cost[v]
		}
	}

	return minV
}
