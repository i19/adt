package booking

import (
	"fmt"
	"testing"
)

func Test_q1(t *testing.T) {
	cap := []int{2, 300, 1}
	price := []int{10, 5, 6}
	count := 302
	// cap := []int{1, 2, 3}
	// price := []int{10, 18, 24}
	// count := 1
	// fmt.Println(q1(cap, price, count))
	fmt.Println(q2(cap, price, count))
	fmt.Println(q2Copy(cap, price, count))
	fmt.Println(q3(cap, price, count))
	fmt.Println(q4(cap, price, count))
	// fmt.Println(minCostDP(cap, price, count))
}
