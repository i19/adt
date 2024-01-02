package leetcode

// 2032. 至少在两个数组中出现的值
// https://leetcode.cn/problems/two-out-of-three/
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	m0, m1, m2, m3 := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	helper := func(nums []int, m map[int]int) {
		for _, num := range nums {
			m0[num] = 1
			m[num] = 1
		}
	}
	helper(nums1, m1)
	helper(nums2, m2)
	helper(nums3, m3)
	for k := range m0 {
		if m1[k]+m2[k]+m3[k] > 1 {
			ans = append(ans, k)
		}
	}
	return
}
