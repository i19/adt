package leetcode

// 189. 轮转数组
// https://leetcode.cn/problems/rotate-array
// O(n) 的空间复杂度
func rotateI(nums []int, k int) {
	k %= len(nums)
	x := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		x[i] = nums[(i-k+len(nums))%len(nums)]
	}
	//
	for i := 0; i < len(nums); i++ {
		nums[i] = x[i]
	}
}

// O(1) 的空间复杂度
func rotateII(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--

	}
}
