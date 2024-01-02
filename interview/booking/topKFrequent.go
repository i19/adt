package booking

import (
	"container/heap"
)

// LCR 060. 前 K 个高频元素
// https://leetcode.cn/problems/g5c51o/description
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	h := &topKFrequentHeap{}
	for n, f := range m {
		heap.Push(h, [2]int{n, f})
		for h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type topKFrequentHeap [][2]int

func (h topKFrequentHeap) Len() int            { return len(h) }
func (h topKFrequentHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h topKFrequentHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h *topKFrequentHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *topKFrequentHeap) Pop() interface{} {
	ans := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ans
}
