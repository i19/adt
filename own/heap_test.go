package own

import "testing"

func TestMaxHeap(t *testing.T) {
	heap := newMaxHeap()

	heap.insert(3)
	heap.insert(1)
	heap.insert(4)
	heap.insert(1)
	heap.insert(5)
	heap.insert(9)
	heap.insert(2)
	heap.insert(6)
	heap.insert(5)

	println("Max Heap:")
	for heap.size > 0 {
		println("%d ", heap.extractMax())
	}

	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	h := initMaxHeapFromData(data)
	println("Max Heap:")
	for h.size > 0 {
		println("%d ", h.extractMax())
	}
}
