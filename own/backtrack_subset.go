package own

// 78. 子集
// https://leetcode.cn/problems/subsets
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 子集型回溯问题的两种写法(或者说代码结构、模板)

// 站在输入角度，对元素的动作有两种，选择或者不选
func subsetsI(nums []int) (res [][]int) {
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		// 不选择该元素
		dfs(i + 1)
		// 选择该元素
		path = append(path, nums[i])
		dfs(i + 1)
		// 针对选择进行回溯
		path = path[:len(path)-1]
	}
	dfs(0)
	return
}

// 站在答案角度，每次必须选择一个元素，第一个选谁，第二个选谁，每个节点都是答案
// 通过规定一个顺序，来避免重复子集的产生（枚举下一项的下标 j > i）
func subsetsII(nums []int) (res [][]int) {
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		// 每个节点都是答案，需要进行保存
		res = append(res, append([]int{}, path...))
		for j := i; j < len(nums); j++ {
			// 选择该元素
			path = append(path, nums[i])
			// 通过下标递增，避免 [1,2] [2,1] 这种情况
			// dfs(i) -> dfs(i+1),dfs(i+2)...dfs(n)
			dfs(j + 1)
			// 针对选择进行回溯
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
