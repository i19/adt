package leetcode

import "sort"

// 4 1 4 , 如果不排序，只去重，那么决策树如下
//
//           ---.--- （1）
//	       /    |
//	  (2) 4  (6)1
//	    / \    /
// (3) 1   4  4(7)
//    /
//(4)4
// 可以看到重复的 [4,1] 和 [1, 4]。 分别是在 两条路径下的 3 和 6 出现的，使用了 used 结构，因为不是跨树传递的所以无法去除
// used 结构是为了避免从一个节点下探的时候，做出重复选择

// 如果增加了排序后得到的就是如下了
//           ---.--- （1）
//	       /    |
//	  (2) 1  (6)4
//	    /     /
// (3) 4     4(7)
//    /
//(4)4

// 因此需要配合排序使用
// 90. 子集 II
// https://leetcode.cn/problems/subsets-ii
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		res = append(res, append([]int{}, path...))
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			// 避免重复选择相同的元素
			if !used[nums[i]] {
				used[nums[i]] = true

				backtrack(i+1, append(path, nums[i]))
			}
			// 下面这是另一种写法， 避免了 map 的使用，同时更充分利用了 排序 的特性
			// if i > start && nums[i] == nums[i-1] {
			// 	continue
			// }
			// backtrack(i+1, append(path, nums[i]))

		}
	}

	backtrack(0, []int{})
	return res
}
