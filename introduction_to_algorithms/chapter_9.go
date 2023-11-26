package introduction_to_algorithms

import (
	"math/rand"
)

func getMinimum(in []int) int {
	x := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] < x {
			x = in[i]
		}
	}
	return x
}

func getMax(in []int) int {
	x := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] > x {
			x = in[i]
		}
	}
	return x
}

func getMinAndMax(in []int) []int {
	min := in[0]
	max := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		} else if in[i] < min {
			min = in[i]
		}
	}
	return []int{min, max}
}

// 一般是用来分析 中位值， 下面这个就是 随机选择 算法
// Yes, there are several algorithms to get the k largest number in an unsorted array.
// One of the algorithms is Quickselect algorithm which is used to find the kth largest element in an unsorted list.
// The algorithm works by selecting a pivot element from the list and partitioning the other elements into two sub-lists,
// according to whether they are less than or greater than the pivot. The sub-list that contains the desired element is then
// recursively sorted until it is of size k.
// The average time complexity of this algorithm is O(n) but it can be O(n^2) in worst case scenarios.

// Another algorithm is to use a max heap data structure.
// We can create a max heap of size k and insert the first k elements of the array into it.
// Then we traverse through the remaining elements of the array and compare each element with the root of the heap.
// If the element is smaller than the root then we skip it,
// otherwise we replace the root with this element and call heapify on the root.
// The time complexity of this algorithm is O(n log k)
// randomizedSelect 是 Quickselect 实现
func randomizedSelect(in []int, start, end, k int) int {
	if start == end {
		return in[start]
	}

	// 随机
	pivot := partitionWithRandom(in, start, end)
	if pivot == k {
		return in[pivot]
	} else if pivot < k {
		return randomizedSelect(in, pivot+1, end, k)
	} else {
		return randomizedSelect(in, start, pivot-1, k)
	}
}

// 下面是 QuickSelect 的另一种实现方法
func quickSelect(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	pivot := arr[rand.Intn(len(arr))]

	var lows, highs, pivots []int
	for _, n := range arr {
		switch {
		case n < pivot:
			lows = append(lows, n)
		case n > pivot:
			highs = append(highs, n)
		default:
			pivots = append(pivots, n)
		}
	}

	switch {
	case k < len(lows):
		return quickSelect(lows, k)
	case k < len(lows)+len(pivots):
		return pivots
	default:
		return quickSelect(highs, k-len(lows)-len(pivots))
	}
}

func E() {

}
