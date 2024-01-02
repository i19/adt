package own

import (
	"fmt"
	"testing"
)

func Test_zero_one_knapsack(t *testing.T) {
	cs := []int{10, 20, 30, 40, 50}
	ws := []int{50, 120, 150, 210, 240}
	c := 50
	// cs := []int{1, 2, 3}
	// ws := []int{5, 11, 15}
	// c := 4
	fmt.Println(dp01pack(cs, ws, c))
	fmt.Println(dp01packII(cs, ws, c))

	fmt.Println(zero_one_knapsack(cs, ws, c))
	fmt.Println(zero_one_knapsack_memo(cs, ws, c))
	fmt.Println(zero_one_knapsack_DP(cs, ws, c))
	fmt.Println(zero_one_knapsack_DP_space(cs, ws, c))
}
