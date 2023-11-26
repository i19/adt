package introduction_to_algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Test_minimum(t *testing.T) {
	for i := 1; i < 100; i++ {
		testIn := rand.Perm(i)
		min := getMinimum(testIn)
		sort.Slice(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		})
		if min != testIn[0] {
			t.Fail()
		}
	}
}

func Test_max(t *testing.T) {
	for i := 1; i < 100; i++ {
		testIn := rand.Perm(i)
		min := getMax(testIn)
		sort.Slice(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		})
		if min != testIn[i-1] {
			t.Fail()
		}
	}
}

func Test_getMinAndMax(t *testing.T) {
	for i := 1; i < 100; i++ {
		testIn := rand.Perm(i)
		res := getMinAndMax(testIn)
		sort.Slice(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		})
		if res[0] != testIn[0] && res[1] != testIn[i-1] {
			t.Fail()
		}
	}
}

func Test_randomizedSelect(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 100; i++ {
		testIn := rand.Perm(i)
		pos := i / 2
		res := randomizedSelect(testIn, 0, i-1, pos)
		sort.Slice(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		})
		if res != testIn[pos] {
			t.Fail()
		}
	}
}

func Test_quickSelect(t *testing.T) {
	arr := []int{3, 7, 2, 1, 8, 5}
	k := 3
	fmt.Println(quickSelect(arr, k))
}
