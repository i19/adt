package introduction_to_algorithms

import (
	"fmt"
	"testing"
	"time"
)

func Test_Fib(t *testing.T) {
	n := 30

	x1 := time.Now()
	//println(FibBottom(n))
	//println(time.Since(x1).Nanoseconds())

	x1 = time.Now()
	println(FibWithMemo(n))
	println(time.Since(x1).Nanoseconds())

	//x1 = time.Now()
	//println(FibRaw(n))
	//println(time.Since(x1).Nanoseconds())
}

func Test_findLCS(t *testing.T) {
	X := "AGGTAB"
	Y := "GXTXAYB"

	lcs := findLCS(X, Y)
	fmt.Println("Longest Common Subsequence:", lcs)
}

func Test_findIncreaseLCS(t *testing.T) {
	fmt.Println(findIncreaseLCS([]int{5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 1, 0, 50, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func Test_longestIncreasingSubsequence(t *testing.T) {
	sequence := []int{5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 1, 0, 50, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	longestIncreasingSubsequence := findLongestIncreasingSubsequence(sequence)
	fmt.Println("Longest Increasing Subsequence:", longestIncreasingSubsequence)
}
