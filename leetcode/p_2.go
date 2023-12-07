package leetcode

// 2. 两数相加
// https://leetcode.cn/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	dummy := head
	step := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			step += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			step += l2.Val
			l2 = l2.Next
		}
		dummy.Next = &ListNode{
			Val: step % 10,
		}
		step /= 10
		dummy = dummy.Next
	}

	return head.Next
}
