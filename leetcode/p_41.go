package leetcode

import "fmt"

// 41. 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
// 这个是我写的解法
// 问题包含，无法正确处理位置正确的（此时 index 不再前进）
func firstMissingPositive(nums []int) int {
	// 方法 1 ，使用 map ，然后从 1 遍历，算法复杂度为 O(n)，空间复杂度为O(n)
	// 方法 2 ，从 1 遍历，判断数组中是否存在，算法复杂度 O(n^2),空间复杂度为（O(1)）
	// 下面结合两种方法的优势，直接在 nums 做原地 哈希

	currentIndex := 0
	currentValue := nums[currentIndex]
	for j := 0; j < len(nums); j++ {
		nextIndex := currentValue % len(nums)
		// if nextIndex == currentIndex {
		// 	nextIndex = (nextIndex + 1) % len(nums)
		// 	continue
		// }
		if nextIndex < 0 {
			nextIndex = -nextIndex
		}
		// println("set index ", nextIndex)
		nextValue := nums[nextIndex]
		nums[nextIndex] = currentValue
		currentValue = nextValue
	}

	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	if nums[0] != len(nums) {
		return len(nums)
	} else {
		return len(nums) + 1
	}
}

func firstMissingPositiveII(nums []int) int {
	// 方法 1 ，使用 map ，然后从 1 遍历，算法复杂度为 O(n)，空间复杂度为O(n)
	// 方法 2 ，从 1 遍历，判断数组中是否存在，算法复杂度 O(n^2),空间复杂度为（O(1)）
	// 下面结合两种方法的优势，直接在 nums 做原地 哈希

	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 && nums[i] <= n {

		}
	}

	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	if nums[0] != len(nums) {
		return len(nums)
	} else {
		return len(nums) + 1
	}
}

// 正确解法(同样是置换的思路)
func firstMissingPositiveIII(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 结果只能是 1 - n+1 中的一个数，因此可以执行前面两个过滤
		// 置换可实现，也是基于这个前提
		// 第二个 for 循环的目的是确保交换后的元素也满足要求。
		// 如果使用 if 语句，只会检查当前元素是否满足要求，而不会检查交换后的元素。
		// 因此，需要使用 for 循环来重复检查交换后的元素，直到它们满足要求为止。
		// 虽然外层有 for ，但是因为下面这个 for 每次都会让一个值放在一个正确位置，所以外层的 for 有可能空转，最后的复杂度仍旧是 O(n)
		for nums[i] >= 1 && nums[i] <= n && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
