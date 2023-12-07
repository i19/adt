package leetcode

// 146. LRU 缓存
// https://leetcode.cn/problems/lru-cache

type doubleListNode struct {
	k    int
	v    int
	Next *doubleListNode
	Prev *doubleListNode
}
type LRUCache struct {
	cache map[int]*doubleListNode
	cap   int
	size  int
	list  *doubleListNode
}

func Constructor(capacity int) LRUCache {
	dummyNode := &doubleListNode{}
	dummyNode.Prev = dummyNode
	dummyNode.Next = dummyNode
	return LRUCache{
		cache: make(map[int]*doubleListNode),
		cap:   capacity,
		size:  0,
		list:  dummyNode,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.v

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.v = value
		this.moveToHead(node)
		return
	}
	if this.isFull() {
		this.size--
		delete(this.cache, this.removeTail().k)
	}
	node := &doubleListNode{k: key, v: value}
	this.cache[key] = node
	this.moveToHead(node)
	this.size++
}

func (this *LRUCache) moveToHead(newHead *doubleListNode) {
	// 如果不为空，则是一个链的内部移动
	if newHead.Prev != nil {
		newHead.Prev.Next = newHead.Next
		newHead.Next.Prev = newHead.Prev
	}

	originHead := this.list.Next
	newHead.Next = originHead
	originHead.Prev = newHead

	newHead.Prev = this.list
	this.list.Next = newHead

	//println("move to head ", key)
}
func (this *LRUCache) removeTail() *doubleListNode {
	lastNode := this.list.Prev
	//println("remove tail start...", k)
	this.list.Prev = lastNode.Prev
	lastNode.Prev.Next = this.list
	//println("remove tail finished...", k)
	return lastNode
}
func (this *LRUCache) isFull() bool {
	return this.size == this.cap
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
