package leetcode

// 746. 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs/?envType=study-plan-v2&envId=dynamic-programming
// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。
//
// 示例 1：
//
// 输入：cost = [10,15,20]
// 输出：15
// 解释：你将从下标为 1 的台阶开始。
// - 支付 15 ，向上爬两个台阶，到达楼梯顶部。
// 总花费为 15 。
// 示例 2：
//
// 输入：cost = [1,100,1,1,1,100,1,1,100,1]
// 输出：6
// 解释：你将从下标为 0 的台阶开始。
// - 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
// - 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
// - 支付 1 ，向上爬一个台阶，到达楼梯顶部。
// 总花费为 6 。
// cost 长度 >= 2
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999

// 这到底应该是一个 DP 问题了
//
//	func minCostClimbingStairs(cost []int) int {
//		dp := make([]int, len(cost))
//		dp[0] = cost[0]
//		dp[1] = cost[1]
//
//		for j := 2; j < len(cost); j++ {
//			dp[j] = min(dp[j-1]+cost[j], dp[j-2]+cost[j])
//		}
//		fmt.Println(dp)
//		return min(dp[len(cost)-1], dp[len(cost)-2])
//	}
func minCostClimbingStairs(cost []int) int {
	pprev := cost[0]
	prev := min(cost[1], cost[0])
	r := 1000

	for i := 2; i < len(cost); i++ {
		r = min(pprev+cost[i], prev)
		pprev = prev
		prev = r
		println("hello world")
	}
	return min(prev, pprev)
}
