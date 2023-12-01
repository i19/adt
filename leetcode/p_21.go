package leetcode

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}
	if list1 != nil {
		dummy.Next = list1
	} else {
		dummy.Next = list2
	}
	return head.Next
}
