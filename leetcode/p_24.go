package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{}

	for prev, first := dummyHead, head; first != nil && first.Next != nil; {
		third := first.Next.Next
		first.Next.Next = first
		prev.Next = first.Next
		first.Next = third
		prev = first
		first = third
	}

	return dummyHead.Next
}
