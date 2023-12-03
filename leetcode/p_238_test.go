package leetcode

import (
	"fmt"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	nums := []int{1, 0, 3, 4}

	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelfII(nums))
}
