package introduction_to_algorithms

import "math/rand"

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
