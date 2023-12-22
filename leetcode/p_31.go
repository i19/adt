package leetcode

// 31. 下一个排列
// https://leetcode.cn/problems/next-permutation
func nextPermutation(nums []int) {
	p := len(nums) - 2
	for ; p >= 0 && nums[p] >= nums[p+1]; p-- {
	}

	if p >= 0 {
		for i := len(nums) - 1; i > p; i-- {
			if nums[i] > nums[p] {
				nums[i], nums[p] = nums[p], nums[i]
				break
			}
		}
	}

	for i := 1; ; i++ {
		l, r := p+i, len(nums)-i
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			break
		}
	}
}
