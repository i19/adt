package leetcode

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
// 栈
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     s :=[]*ListNode{&ListNode{Next:head}}
//     for head != nil {
//         s = append(s, head)
//         head = head.Next
//     }

//	    s[len(s)-1-n].Next = s[len(s)-1-n].Next.Next
//	    return s[0].Next
//	}
//
// 双指针
func removeNthFromEndDoublePointer(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	for i := 0; i < n; i++ {
		second = second.Next
	}

	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}
