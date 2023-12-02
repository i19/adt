package leetcode

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum
import "container/heap"

var duplicate []int

type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i int, j int) bool {
	return duplicate[m[i]] > duplicate[m[j]]
}

func (m myHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *myHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	duplicate = nums
	var res []int

	var x myHeap
	heap.Init(&x)
	for i := 0; i < len(nums); i++ {
		heap.Push(&x, i)
		if i+1 >= k {
			for {
				maxI := x[0]
				if maxI >= i-k+1 {
					res = append(res, nums[maxI])
					break
				}
				heap.Pop(&x)
			}
		}
	}

	return res
}
