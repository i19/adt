package leetcode

// 740. 删除并获得点数
// https://leetcode.cn/problems/delete-and-earn/?envType=study-plan-v2&envId=dynamic-programming
// 给你一个整数数组 nums ，你可以对它进行一些操作。
//
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
//
// 示例 1：
//
// 输入：nums = [3,4,2]
// 输出：6
// 解释：
// 删除 4 获得 4 个点数，因此 3 也被删除。
// 之后，删除 2 获得 2 个点数。总共获得 6 个点数。
// 示例 2：
//
// 输入：nums = [2,2,3,3,3,4]
// 输出：9
// 解释：
// 删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
// 之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
// 总共获得 9 个点数。
//
// 提示：
//
// 1 <= nums.length <= 2 * 104
// 1 <= nums[i] <= 104
// 打家劫舍问题变种，但是感觉有点太硬了
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	sums := make([]int, maxVal+1)
	for _, num := range nums {
		sums[num] += num
	}

	if len(sums) == 1 {
		return sums[0]
	}
	dp := make([]int, len(sums))
	dp[0] = sums[0]
	dp[1] = max(sums[0], sums[1])
	for i := 2; i < len(sums); i++ {
		dp[i] = max(dp[i-2]+sums[i], dp[i-1])
	}

	return max(dp[len(sums)-1], dp[len(sums)-2])
}
