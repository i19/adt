package leetcode

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数，
// 子数组是数组中元素的【连续非空序列】。
// 前缀和的思路还是蛮有趣的
// 注意这里的开闭空间，都是开空间
// pre[i] = sum(0..i)
// 所以 sum[i..j] = pre[j] - pre[i-1]
// 下面的代码，则是借鉴了这个思路
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	count, pre := 0, 0
	// 如果没有下面这个赋值，那么就会漏掉第一次相等的情况
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre] += 1
	}
	return count
}
