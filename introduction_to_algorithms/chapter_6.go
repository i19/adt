package introduction_to_algorithms

// HEAP SORTING
// 大根堆，小根堆的性质 （完全二叉树）
// 1. 给定节点 i，其 PARENT(i) = i/2 LEFT(i) = i*2 + 1 RIGHT(i) = i*2+2
// 2. 大根堆，孩子节点小于父节点，根节点最大。 小根堆反之
// heapify 的目的就是维持堆的性质，使之满足堆的性质
// 经过梳理，可以理解 heapify 的时间复杂度与 以指定节点为根的子树高度有关
// 也就是 heapify 的复杂度是 O(h) 而 h的高度为 log_n
// 所以 heapify 的复杂度为 O(log_n)
func heapify(arr []int, heapSize, nodeOffset int) {
	leftChildOffset := nodeOffset*2 + 1
	rightChildOffset := nodeOffset*2 + 2
	var largeOffset int
	if leftChildOffset < heapSize && arr[leftChildOffset] > arr[nodeOffset] {
		largeOffset = leftChildOffset
	} else {
		largeOffset = nodeOffset
	}

	if rightChildOffset < heapSize && arr[rightChildOffset] > arr[largeOffset] {
		largeOffset = rightChildOffset
	}

	if largeOffset != nodeOffset {
		arr[nodeOffset], arr[largeOffset] = arr[largeOffset], arr[nodeOffset]
		heapify(arr, heapSize, largeOffset)
	}
}

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
func buildHeap(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
}

// 将最大的放到最后
// 然后通过 heapify 的 heapSize 参数使后面的不参与 堆化
// 通过每次 heapify 使最大的放在 arr[0] 换走，再堆化
// 最终完成排序
// O(nlog_n)
func heapSort(arr []int) {
	buildHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- { //O(n)
		arr[0], arr[i] = arr[i], arr[0] // O(1)
		heapify(arr, i, 0)              // O(log_n)
	}
}
