package leetcode

// 4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	s1, s2 := 0, 0
	i, j, k := 0, 0, 0
	for ; k < l/2+1 && i < len(nums1) && j < len(nums2); k++ {
		s1 = s2
		if nums1[i] >= nums2[j] {
			s2 = nums2[j]
			j++
		} else {
			s2 = nums1[i]
			i++
		}
	}
	for ; k < l/2+1 && i < len(nums1); i++ {
		s1 = s2
		s2 = nums1[i]
		k++
	}

	for ; k < l/2+1 && j < len(nums2); j++ {
		s1 = s2
		s2 = nums2[j]
		k++
	}

	if l%2 == 1 {
		return float64(s2)
	} else {
		return float64(s1+s2) / 2
	}
}
