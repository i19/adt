package own

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsI(nums))
	fmt.Println(subsetsII(nums))
}
