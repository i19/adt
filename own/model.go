package own

type linkListNode struct {
	value interface{}
	next  *linkListNode
}

type hashPair struct {
	key int
	val string
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}
