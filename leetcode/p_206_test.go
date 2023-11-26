package leetcode

import (
	"encoding/json"
	"testing"
)

func Test_reverseList(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	r := &ListNode{}
	copyR := r
	for i := 0; i < len(input)-1; i++ {
		r.Val = input[i]
		r.Next = &ListNode{}
		r = r.Next
	}
	r.Val = input[len(input)-1]

	r = reverseListInRecursion(copyR)
	zz, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	println(string(zz))
}
