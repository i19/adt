package leetcode

// https://leetcode.cn/problems/reverse-linked-list/
// 206. 反转链表
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历解法
func reverseListInTraversing(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // 保存指针
		curr.Next = prev  // 指针反向
		prev = curr       // 递进
		curr = next       // 递进
	}
	return prev
}

// 递归解法
func reverseListInRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListInRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
