package introduction_to_algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_quicksort(t *testing.T) {
	for i := 0; i < 100; i++ {
		//i := 10
		testIn := rand.Perm(i)
		//testIn := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
		quickSort(testIn, 0, len(testIn)-1)
		if !sort.SliceIsSorted(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		}) {
			t.Fail()
		}
		fmt.Println(testIn)
	}
}
