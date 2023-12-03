package leetcode

// 238. 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self
func productExceptSelf(nums []int) []int {
	var res []int
	for skip := 0; skip < len(nums); skip++ {
		tr := 1
		for noSkip := 0; noSkip < len(nums); noSkip++ {
			if noSkip != skip {
				tr *= nums[noSkip]
			}
		}
		res = append(res, tr)
	}
	return res
}

// 扩展前缀和，使用一个前缀一个后缀，分别维护乘积
// pre[i] = 1*2*3*..*i
// aft[i] = len(num)-1 * len(num)-2 *len(num)-3 * ... * i
func productExceptSelfII(nums []int) []int {
	pre := make([]int, len(nums))
	aft := make([]int, len(nums))
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}

	aft[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		aft[i] = nums[i+1] * aft[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = pre[i] * aft[i]
	}
	return res
}
