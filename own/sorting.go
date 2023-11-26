package own

import "math/rand"

// 选择排序 O(n^2)
func selectSort(in []int) {
	if len(in) <= 1 {
		return
	}

	for i := 0; i < len(in); i++ {
		min := in[i]
		minIndex := i
		for j := i + 1; j < len(in); j++ {
			if in[j] < min {
				min = in[j]
				minIndex = j
			}
		}
		in[i], in[minIndex] = in[minIndex], in[i]
	}
}

// 冒泡排序 O(n^2)
func bubbleSort(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 插入排序 O(n^2)
func insertSort(in []int) {
	if len(in) <= 1 {
		return
	}

	for j := 1; j < len(in); j++ {
		key := in[j]
		//Insert A[j] into the sorted sequence A[0.. j-1]
		i := j - 1
		for ; i >= 0 && in[i] > key; i-- {
			in[i+1] = in[i]
		}
		in[i+1] = key
	}
}

// 插入排序递归实现
func insertSortUsingRecursion(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	helper := func(arr []int, val int) []int {
		arr = append(arr, val)
		for i := len(arr) - 2; i >= 0; i-- {
			if arr[i] > val {
				arr[i+1] = arr[i]
				if i == 0 {
					arr[i] = val
				}
			} else {
				arr[i+1] = val
				break
			}
		}
		return arr
	}
	return helper(insertSortUsingRecursion(in[:len(in)-1]), in[len(in)-1])
}

/* 计数排序 */
// 简单实现，无法用于排序对象
func countingSortNaive(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 遍历 counter ，将各元素填入原数组 nums
	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

// mergeSort 是一种分治 nlog_n
func mergeSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	mergeHelper := func(left, right []int) []int {
		var result []int

		l, r := 0, 0
		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				result = append(result, left[l])
				l++
			} else {
				result = append(result, right[r])
				r++
			}
		}

		result = append(result, left[l:]...)
		result = append(result, right[r:]...)

		return result
	}

	mid := len(in) / 2
	left := mergeSort(in[:mid])
	right := mergeSort(in[mid:])
	return mergeHelper(left, right)
}

// 快排是一种分治的思路，有 2 个性质
// 1. 基准值之前的要比基准值小
// 2. 基准值之后的要比基准值大
// 然后将 基准值前后（不包含基准值）的数组再进行排序（分治）
// 复杂度分析
// 最坏的情况是 pivot 后，两边的比例是 1 ：N 这样，就是 O(n^2)
// 最好情况是，两边的比例是 1:1 也就是均衡，这样就是 O(nlg_n)
func quickSort(arr []int, begin, end int) {
	if begin < end {
		pivot := partition1(arr, begin, end)
		//pivot := partition2(arr, begin, end)
		//pivot := partitionWithRandom(arr, begin, end)
		quickSort(arr, begin, pivot-1)
		quickSort(arr, pivot+1, end)
	}
}

// partition 的功能就是，就是为了满足上面的 性质，可以使用快慢指针实现
// 具体实现方式：
// 选择一个基准数，放在合适的位置
// 基准数的选择，一般 是 begin 或者 end 位置的值，如果选择其他位置，则需要多一次 比较操作（）
// 利用快慢指针实现，慢指针（包含）之前的值 要比基准值小，慢指针之后的要比基准值 大
// 最后将基准值和慢指针后面的位置进行交换，同时将该位置返回给 quickSort 进行分治
// 此时，就已经找到了 基准值 在数组中所在的位置了，那么后面只需要将 该位置 前后（不包含）的进行分治即可，所以最后返回的是基准值的最终位置
func partition1(arr []int, begin, end int) int {
	pivot := arr[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}

// 也可以使用对向双指针实现
func partition2(arr []int, begin, end int) int {
	pivot := arr[begin]
	for j := begin + 1; j <= end; {
		if arr[j] <= pivot {
			j++
		} else {
			arr[j], arr[end] = arr[end], arr[j]
			end--
		}
	}
	arr[begin], arr[end] = arr[end], arr[begin]
	return end
}

// 随机策略，平均复杂度稳定
func partitionWithRandom(arr []int, begin, end int) int {
	randPos := begin + rand.Intn(end-begin)
	arr[begin], arr[randPos] = arr[randPos], arr[begin]

	pivot := arr[begin]
	for j := begin + 1; j <= end; {
		if arr[j] <= pivot {
			j++
		} else {
			arr[j], arr[end] = arr[end], arr[j]
			end--
		}
	}
	arr[begin], arr[end] = arr[end], arr[begin]
	return end
}

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
