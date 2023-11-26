package own

// 大根堆，小根堆的性质 （完全二叉树）
// 1. 给定节点 i，其 PARENT(i) = (i-1)/2 LEFT(i) = i*2 + 1 RIGHT(i) = i*2+2
// 2. 大根堆，孩子节点小于父节点，根节点最大。 小根堆反之
// heapify 的目的就是维持堆的性质，使之满足堆的性质
// 经过梳理，可以理解 heapify 的时间复杂度与 以指定节点为根的子树高度有关
// 也就是 heapify 的复杂度是 O(h) 而 h的高度为 log_n
// 所以 heapify 的复杂度为 O(log_n)

// buildHeap 自下向上建堆。（将非叶子节点执行 heapify 操作）
// 堆是完全二叉树，平铺展开在数组 A 中， 那么数组后面 A[n/2:] 将都是叶子节点
// 那么 A[n/2 - 1] 开始向前遍历到 0 （根节点），执行 heapify 操作，将获取一个堆结构的数组
// 这其实是一种优化后的建堆方法，从倒数第二层的最后一个节点开始
// 您提到的算法导论中提到的是一种优化的建堆方法，称为“线性建堆”（linear-time heap construction）。
// 这种方法的时间复杂度确实是O(n)。 线性建堆的思想是从倒数第二层开始，依次向上进行堆化。
// 具体来说，从倒数第二层的最后一个节点开始，依次向前遍历每个非叶子节点，对每个节点进行向下堆化操作。
// 由于叶子节点已经满足堆的条件，所以只需要对非叶子节点进行堆化操作即可。因为倒数第二层最多有n/4个节点，倒数第三层最多有n/8个节点，
// 以此类推，所以总共需要进行O(n)次堆化操作，时间复杂度为O(n)。
// 需要注意的是，虽然线性建堆的时间复杂度比传统的建堆方法要优秀，
// 但是它的常数因子比较大，所以在实际应用中，并不一定比传统的建堆方法更快。
// 同时，线性建堆的实现也比较复杂，需要对堆化操作进行优化，
// 以及对不同层次的节点进行分类处理等等。因此，在实际应用中，需要根据具体情况选择合适的建堆方法。

type MaxHeap struct {
	array []int
	size  int
}

// 初始化一个最大堆
func newMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 通过已有数据初始化一个最大堆
// 该过程的 时间复杂度为 O(n) 非常搞笑
func initMaxHeapFromData(data []int) *MaxHeap {
	r := &MaxHeap{array: data, size: len(data)}
	for i := r.size/2 - 1; i >= 0; i-- {
		r.heapifyDown(i)
	}
	return r
}

// 上移操作，保持堆的性质
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index] > h.array[parentIndex] {
			h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
			index = parentIndex
		} else {
			break
		}
	}
}

// 下移操作，保持堆的性质
func (h *MaxHeap) heapifyDown(index int) {
	for index*2+1 < h.size {
		leftChild := index*2 + 1
		rightChild := index*2 + 2
		largerChildIndex := leftChild

		if rightChild < h.size && h.array[rightChild] > h.array[leftChild] {
			largerChildIndex = rightChild
		}

		if h.array[index] < h.array[largerChildIndex] {
			h.array[index], h.array[largerChildIndex] = h.array[largerChildIndex], h.array[index]
			index = largerChildIndex
		} else {
			break
		}
	}
}

// 插入元素
func (h *MaxHeap) insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp(h.size - 1)
}

// 弹出堆顶元素
// 堆顶元素是二叉树的根节点，即列表首元素。如果我们直接从列表中删除首元素，那么二叉树中所有节点的索引都会发生变化，
// 这将使得后续使用堆化修复变得困难。为了尽量减少元素索引的变动，我们采用以下操作步骤。
//  1. 暂存堆顶。并将堆底移至堆顶
//  2. 将堆底从列表中删除（size--）
//  3. 从根节点开始，从顶至底执行堆化。
func (h *MaxHeap) extractMax() int {
	if h.size == 0 {
		return -1
	}

	max := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--
	h.heapifyDown(0)

	return max
}

// 获取堆顶元素
func (h *MaxHeap) getMax() int {
	if h.size == 0 {
		return -1
	}
	return h.array[0]
}
