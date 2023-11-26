package own

import "testing"

func Test_binarySearch(t *testing.T) {
	println(binarySearch([]int{1, 2, 3, 4, 5}, 4))
	println(binarySearch([]int{1, 2, 3, 4, 5}, 5))

	repeatTargets := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5}
	println(binaryFindLeft(repeatTargets, 6), binaryFindRight(repeatTargets, 6))
}
