package main

import (
	"container/heap"
	"fmt"
)

var (
	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
)

type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i int, j int) bool {
	return nums[m[i]] > nums[m[j]]
}

func (m myHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m myHeap) Pop() any {
	panic("not implemented") // TODO: Implement
}

func main() {
	var x myHeap
	for i := 0; i < len(nums); i++ {
		heap.Push(&x, i)
	}
	fmt.Println(x[0])
}
