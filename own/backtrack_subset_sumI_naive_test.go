package own

import (
	"fmt"
	"testing"
)

func Test_ssnI(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	fmt.Println(ssnI(nums, target))
}

func Test_ssnII(t *testing.T) {
	nums := []int{4, 4, 5}
	target := 9
	fmt.Println(ssnI(nums, target))
	fmt.Println(ssnII(nums, target))
}
