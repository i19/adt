package leetcode

// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures
// 单调栈机制
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {

		for len(stack) != 0 {
			peakV := stack[len(stack)-1]
			if temperatures[peakV] < temperatures[i] {
				res[peakV] = i - peakV
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}
