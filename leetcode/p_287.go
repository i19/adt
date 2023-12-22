package leetcode

// 287. 寻找重复数
// https://leetcode.cn/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return 0
}

func findDuplicateII(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
