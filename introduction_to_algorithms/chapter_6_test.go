package introduction_to_algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func checkIsHeap(arr []int) bool {
	for i := 0; i < len(arr)/2-1; i++ {
		left := i*2 + 1
		right := left + 1
		if arr[i] > arr[left] && arr[i] > arr[right] {
			continue
		} else {
			return false
		}
	}
	return true
}

func Test_buildHeap(t *testing.T) {
	x := rand.Perm(100)
	buildHeap(x)
	fmt.Print(x, "\n")
	fmt.Print(checkIsHeap(x))
}

func Test_heapSort(t *testing.T) {
	for i := 0; i < 1; i++ {
		testIn := rand.Perm(i)
		heapSort(testIn)
		if !sort.SliceIsSorted(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		}) {
			t.Fail()
		}
	}
}
