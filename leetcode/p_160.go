package leetcode

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists

// 方法一：按照指针地址寻址，空间复杂度为 O(m+n)
func getIntersectionNodeI(headA, headB *ListNode) *ListNode {
	da, db := headA, headB
	m := make(map[*ListNode]bool)
	for da != nil {
		m[da] = true
		da = da.Next
	}
	for db != nil {
		if m[db] {
			return db
		}
		db = db.Next
	}
	return nil
}

// 方法二：利用双指针，每个指针都走一遍 A 和 B，返回最后那个节点（可能是 nil）
// 想法是这样的，每次大家『都』向前进 1 ，如果为空了，则换到另一条线上。相当于每个指针都走了两条路，这样长度就相同了
// 最后两个指针指向的一定会相等，这个位置要么是空(nil 相等也算，相当于走了 la+lb+1)，要么就是相交的点(nil 相等也算，相当于走了 la+lb)
func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
	da, db := headA, headB
	if da == nil || db == nil {
		return nil
	}
	for da != db {
		if da == nil {
			da = headB
		} else {
			da = da.Next
		}

		if db == nil {
			db = headA
		} else {
			db = db.Next
		}
	}

	return da
}
