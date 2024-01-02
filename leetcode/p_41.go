package leetcode

// 41. 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 结果只能是 [1, n+1] 中的一个数，因此可以执行前面两个过滤
		// 置换可实现，也是基于这个前提
		// 第二个 for 循环的目的是确保交换后的元素也满足要求。
		// 如果使用 if 语句，只会检查当前元素是否满足要求，而不会检查交换后的元素。
		// 因此，需要使用 for 循环来重复检查交换后的元素，直到它们满足要求为止。
		// 虽然外层有 for ，但是因为下面这个 for 每次都会让一个值放在一个正确位置，所以外层的 for 有可能空转，最后的复杂度仍旧是 O(n)
		// leetcode 解法中隐去了 nums[i] != i+1 的判断
		// 因为 nums[i] != nums[nums[i]-1] 隐含了这个判断
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
