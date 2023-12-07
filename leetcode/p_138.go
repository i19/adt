package leetcode

// 138. 随机链表的复制
// https://leetcode.cn/problems/copy-list-with-random-pointer/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m1 := make(map[*Node]int, 0)
	x := []*Node{}
	for i, dummy := 0, head; dummy != nil; i++ {
		x = append(x, &Node{Val: dummy.Val})
		m1[dummy] = i
		dummy = dummy.Next
	}
	for i := 0; i < len(x)-1; i++ {
		x[i].Next = x[i+1]
	}

	// nth -> nth
	m2 := make(map[int]int, 0)
	for i, dummy := 0, head; dummy != nil; i++ {
		if dummy.Random != nil {
			m2[i] = m1[dummy.Random]
		}
		dummy = dummy.Next
	}

	for i := 0; i < len(x); i++ {
		if v, ok := m2[i]; ok {
			x[i].Random = x[v]
		}
	}
	return x[0]
}
