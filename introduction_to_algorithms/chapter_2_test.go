package introduction_to_algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_insertSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIn := rand.Perm(i)
		insertSort(testIn)
		if !sort.SliceIsSorted(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		}) {
			t.Fail()
		}
	}
}

func Test_insertSortUsingRecursion(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIn := rand.Perm(i)
		testOut := insertSortUsingRecursion(testIn)
		if !sort.SliceIsSorted(testOut, func(i, j int) bool {
			return testOut[i] < testOut[j]
		}) || len(testIn) != len(testOut) {
			t.Fail()
		}
	}
}

func Test_selectSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIn := rand.Perm(i)
		selectSort(testIn)
		if !sort.SliceIsSorted(testIn, func(i, j int) bool {
			return testIn[i] < testIn[j]
		}) {
			t.Fail()
		}
	}
}

func Test_mergeSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIn := rand.Perm(i)
		testOut := mergeSort(testIn)
		if !sort.SliceIsSorted(testOut, func(i, j int) bool {
			return testOut[i] < testOut[j]
		}) {
			t.Fatal(testOut)
		}
	}
}

func Test_binarySearch(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 89, 100, 102}
	for i := 0; i < len(x); i++ {
		if i != binarySearch(x, x[i]) {
			t.Fatal()
		}
	}
	if binarySearch(x, 0) != -1 {
		t.Fail()
	}
	if binarySearch(x, 1000) != -1 {
		t.Fail()
	}
}

func Test_twoSum(t *testing.T) {
	//x := []int{}
	x := []int{1, 5, 3, 99, 98, 2}
	println(twoSum(x, 100))
}

func Test_findReversePairV1(t *testing.T) {
	input := []int{5, 4, 3, 2, 1, 0}
	result := findReversePairV1(input)
	fmt.Print(result)
}
func Test_findReversePairV2(t *testing.T) {
	input := []int{5, 4, 3, 2, 1, 0}
	_, count := findReversePairV2(input)
	fmt.Print(count)
}
func Test_findReversePairV3(t *testing.T) {
	input := []int{5, 4, 3, 2, 1, 0}
	_, result := findReversePairV3(input)
	fmt.Print(result)
}
