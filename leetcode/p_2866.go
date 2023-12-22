package leetcode

// 2866. 美丽塔 II
// https://leetcode.cn/problems/beautiful-towers-ii
// 暴力运算
// func maximumSumOfHeights(maxHeights []int) int64 {
//     res := 0
//     for mid:=0; mid < len(maxHeights); mid++{
//         tres := maxHeights[mid]
//         lv,rv := tres, tres
//         for left := mid-1; left >= 0; left-- {
//             lv = min(maxHeights[left], lv)
//             tres += lv
//         }
//         for right := mid+1; right < len(maxHeights); right++ {
//             rv = min(maxHeights[right], rv)
//             tres+=rv
//         }
//         res=max(res, tres)
//     }
//     return int64(res)
// }

// 使用单调栈解决
func maximumSumOfHeights(maxHeights []int) int64 {
	left := make([]int, len(maxHeights))
	right := make([]int, len(maxHeights))
	stackL, stackR := []int{}, []int{}
	for i := 0; i < len(maxHeights); i++ {
		for len(stackL) > 0 && maxHeights[i] < maxHeights[stackL[len(stackL)-1]] {
			stackL = stackL[:len(stackL)-1]
		}
		if len(stackL) == 0 {
			left[i] = (i + 1) * maxHeights[i]
		} else {
			last := stackL[len(stackL)-1]
			left[i] = left[last] + (i-last)*maxHeights[i]
		}
		stackL = append(stackL, i)
	}

	res := 0
	for i := len(maxHeights) - 1; i > -1; i-- {
		for len(stackR) > 0 && maxHeights[i] < maxHeights[stackR[len(stackR)-1]] {
			stackR = stackR[:len(stackR)-1]
		}
		if len(stackR) == 0 {
			right[i] = (len(maxHeights) - i) * maxHeights[i]
		} else {
			last := stackR[len(stackR)-1]
			right[i] = right[last] + (last-i)*maxHeights[i]
		}
		stackR = append(stackR, i)
		res = max(res, left[i]+right[i]-maxHeights[i])
	}
	return int64(res)
}
