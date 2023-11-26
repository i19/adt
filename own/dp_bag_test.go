package own

import (
	"fmt"
	"testing"
)

func Test_dp01pack(t *testing.T) {
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	c := 50
	println(dp01pack(wgt, val, c))
	println(dp01packII(wgt, val, c))

	wgt = []int{1, 2, 3}
	val = []int{5, 11, 15}
	c = 4
	println(dpFullpack(wgt, val, c))
}

func Test_Coins(t *testing.T) {
	coins := []int{1, 2}
	amount := 5

	// coins := []int{2}
	// amount := 3
	// println(coinChangeII(coins, amount))
	println(coinChangIII(coins, amount))
	for _, r := range coinChangIIII(coins, amount) {
		fmt.Println(r)
	}
}
