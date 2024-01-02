package booking

import "math"

// 1334. 阈值距离内邻居最少的城市
// https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	newEdges := make([][]int, n)
	for i := 0; i < n; i++ {
		newEdges[i] = make([]int, n)
	}
	for _, edge := range edges {
		newEdges[edge[0]][edge[1]] = edge[2]
		newEdges[edge[1]][edge[0]] = edge[2]
	}

	selectd := make([]bool, n)
	var dfs func(from, remain int, arrived []bool)
	dfs = func(from, remain int, arrived []bool) {
		if remain <= 0 {
			return
		}

		for i := 0; i < n; i++ {
			cost := newEdges[from][i]
			if !selectd[i] && newEdges[from][i] > 0 && remain >= cost {
				selectd[i] = true
				arrived[i] = true
				dfs(i, remain-cost, arrived)
				selectd[i] = false
			}
		}
	}
	arrivedCity := make([][]int, n)
	for i := 0; i < n; i++ {
		selectd[i] = true
		arrived := make([]bool, n)
		dfs(i, distanceThreshold, arrived)
		cities := []int{}
		for j := 0; j < n; j++ {
			if arrived[j] {
				cities = append(cities, j)
			}
		}
		arrivedCity[i] = cities
		selectd[i] = false
	}
	minI, minH := 0, n
	for i, y := range arrivedCity {
		if len(y) <= minH {
			minH = len(y)
			minI = i
		}
	}
	return minI
}

// Floyd 算法，借助 k 点
func findTheCityII(n int, edges [][]int, distanceThreshold int) int {
	weights := make([][]int, n)
	for i := 0; i < n; i++ {
		weights[i] = make([]int, n)
		for j := 0; j < n; j++ {
			weights[i][j] = math.MaxInt32 / 2
		}
	}
	for _, edge := range edges {
		weights[edge[0]][edge[1]] = edge[2]
		weights[edge[1]][edge[0]] = edge[2]
	}

	var dfs func(k, i, j int) int
	dfs = func(k, i, j int) int {
		if k < 0 {
			return weights[i][j]
		}
		return min(dfs(k-1, i, j), dfs(k-1, i, k)+dfs(k-1, k, j))
	}

	minCnt := n
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i != j {
				dis := dfs(n-1, i, j)
				if dis <= distanceThreshold {
					cnt++
				}
			}
		}
		if cnt <= minCnt {
			minCnt = cnt
			ans = i
		}
	}

	return ans
}

// Floyd 算法，修改成 DP
func findTheCityIIDP(n int, edges [][]int, distanceThreshold int) int {
	weights := make([][]int, n)
	for i := 0; i < n; i++ {
		weights[i] = make([]int, n)
		for j := 0; j < n; j++ {
			weights[i][j] = math.MaxInt32 / 2
		}
	}
	for _, edge := range edges {
		weights[edge[0]][edge[1]] = edge[2]
		weights[edge[1]][edge[0]] = edge[2]
	}

	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	dp[0] = weights
	for k := 1; k <= n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					dp[k][i][j] = min(dp[k-1][i][j], dp[k-1][i][k-1]+dp[k-1][k-1][j])
				}
			}
		}
	}

	ans := 0
	minCnt := n
	for i, dis := range dp[n] {
		cnt := 0
		for j, d := range dis {
			if i != j && d <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= minCnt {
			ans = i
			minCnt = cnt
		}
	}

	return ans
}
