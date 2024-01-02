package booking

import (
	"container/heap"
	"strings"
)

// 2512. 奖励最顶尖的 K 名学生
// https://leetcode.cn/problems/reward-top-k-students/
func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	pm, nm := make(map[string]int), make(map[string]int)
	for _, v := range positive_feedback {
		pm[v] = 3
	}
	for _, v := range negative_feedback {
		nm[v] = -1
	}
	h := &topHeap{}
	for i, id := range student_id {
		x := [2]int{id, 0}
		for _, word := range strings.Split(report[i], " ") {
			x[1] += pm[word]
			x[1] += nm[word]
		}
		println(x[0], x[1])
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 因为是小根堆，而结果需要降序排序，所以这里是逆序填充的
	ans := make([]int, h.Len()-1)
	for h.Len() != 0 {
		ans[h.Len()-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type topHeap [][2]int

func (h topHeap) Len() int      { return len(h) }
func (h topHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 小根堆，通过 pop 去除掉低分高ID
func (h topHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] || (h[i][1] == h[j][1] && h[i][0] > h[j][0])
}
func (h *topHeap) Push(x interface{})     { *h = append(*h, x.([2]int)) }
func (h *topHeap) Pop() (ans interface{}) { ans = (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return }
