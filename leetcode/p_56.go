package leetcode

import (
	"sort"
)

// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals
func mergeI(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i := 0; i < len(intervals); i++ {
		newStart := len(res)
		newEnd := len(res)
		for j := 0; j < len(res); j++ {
			// 找到起始点
			if res[j][0] <= intervals[i][0] && res[j][1] >= intervals[i][0] && j < newStart {
				newStart = j
			}
			// 找到结束点
			if res[j][0] <= intervals[i][1] && res[j][1] >= intervals[i][1] && j < newEnd {
				newEnd = j
			}
		}
		if newStart == len(res) && newEnd == len(res) {
			res = append(res, intervals[i])
		} else if newStart != len(res) && newEnd != len(res) {
			res[newStart][1] = res[newEnd][1]
			res = append(res[:newStart], res[newEnd:]...)
		} else if newStart == len(res) {
			res[newEnd][0] = intervals[i][0]
		} else if newEnd == len(res) {
			res[newStart][1] = intervals[i][1]
		}
	}
	return res
}

// 调优
func mergeII(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	intervals2 := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] != intervals2[len(intervals2)-1][0] {
			intervals2 = append(intervals2, intervals[i])
		} else {
			intervals2[len(intervals2)-1][1] = max(intervals2[len(intervals2)-1][1], intervals[i][1])
		}
	}
	res := [][]int{intervals2[0]}
	for i := 1; i < len(intervals2); i++ {
		if intervals2[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = intervals2[i][1]
		} else {
			res = append(res, intervals2[i])
		}
	}
	return res
}

// 调优
func mergeIII(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}
