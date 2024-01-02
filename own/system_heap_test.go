package own

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap_Heap(t *testing.T) {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                // 初始化heap
	heap.Push(h, 6)                             // 调用push
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
