package leetcode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	tm := swapPairs(head)
	x, _ := json.Marshal(tm)
	fmt.Println(string(x))

}
