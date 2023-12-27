package leetcode

// 1276. 不浪费原料的汉堡制作方案
// https://leetcode.cn/problems/number-of-burgers-with-no-waste-of-ingredients

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	bc := (tomatoSlices - 2*cheeseSlices) / 2
	sc := (4*cheeseSlices - tomatoSlices) / 2
	if bc < 0 || sc < 0 || (bc*4+sc*2) != tomatoSlices || (bc+sc) != cheeseSlices {
		return []int{}
	} else {
		return []int{bc, sc}
	}
}
