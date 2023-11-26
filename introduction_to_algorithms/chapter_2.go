package introduction_to_algorithms

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

// lg_n
func binarySearch(sortedIn []int, value int) int {
	left := 0
	right := len(sortedIn) - 1
	for left <= right {
		mid := (left + right) / 2
		if sortedIn[mid] == value {
			return mid
		} else if sortedIn[mid] > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func twoSum(in []int, x int) bool {
	sortedArr := mergeSort(in)
	for i, j := 0, len(sortedArr)-1; i < j; {
		v := sortedArr[i] + sortedArr[j]
		if v == x {
			return true
		} else if v > x {
			j--
		} else {
			i++
		}
	}

	return false
}

// O(n(n-1)/2)  O(n^2) 返回反序对
func findReversePairV1(in []int) [][]int {
	if len(in) <= 1 {
		return nil
	}

	var result [][]int
	for i := 0; i < len(in)-1; i++ {
		for j := i + 1; j < len(in); j++ {
			if in[j] < in[i] {
				result = append(result, []int{in[i], in[j]})
			}
		}
	}

	return result
}

// O(nlg_n) 使用 mergesort 思路 返回数量
func findReversePairV2(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}

	helper := func(left, right []int) ([]int, int) {
		var result []int
		count := 0
		l, r := 0, 0
		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				result = append(result, left[l])
				l++
			} else {
				result = append(result, right[r])
				r++
				count += len(left) - l
			}
		}

		result = append(result, left[l:]...)
		result = append(result, right[r:]...)
		return result, count
	}

	mid := len(arr) / 2
	left, leftCount := findReversePairV2(arr[:mid])
	right, rightCount := findReversePairV2(arr[mid:])
	result, mergeCount := helper(left, right)
	return result, leftCount + rightCount + mergeCount
}

// O(nlg_n) 使用 mergesort 思路，返回反序对
func findReversePairV3(arr []int) ([]int, [][]int) {
	if len(arr) <= 1 {
		return arr, nil
	}

	helper := func(left, right []int) ([]int, [][]int) {
		var result []int
		var reverse [][]int

		l, r := 0, 0
		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				result = append(result, left[l])
				l++
			} else {
				result = append(result, right[r])
				for x := l; x < len(left); x++ {
					reverse = append(reverse, []int{left[x], right[r]})
				}
				r++
			}
		}

		result = append(result, left[l:]...)
		result = append(result, right[r:]...)
		return result, reverse
	}

	mid := len(arr) / 2
	left, leftReverse := findReversePairV3(arr[:mid])
	right, rightReverse := findReversePairV3(arr[mid:])
	result, mergeReverse := helper(left, right)
	var x [][]int
	x = append(x, leftReverse...)
	x = append(x, rightReverse...)
	x = append(x, mergeReverse...)
	return result, x
}
