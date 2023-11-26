package leetcode

import "testing"

func Test_wordBreak(t *testing.T) {
	println(wordBreak("leetcode", []string{"leet", "code", "abc"}))
}

func Test_longestPalindromeSubseq(t *testing.T) {
	//println(longestPalindromeSubseq("bbbab"))
	println(longestPalindromeSubseq("bab"))
}

func Test_minDistance(t *testing.T) {
	println(minDistance("", "ros"))
	println(minDistance("horse", "ros"))
	println(minDistance("intention", "execution"))
	//intention
	//execution
}

func Test_maximumMinutes(t *testing.T) {
	println(maximumMinutes([][]int{
		{0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 1, 0},
		{0, 2, 0, 0, 1, 2, 0},
		{0, 0, 2, 2, 2, 0, 2},
		{0, 0, 0, 0, 0, 0, 0},
	}))
	println(maximumMinutes([][]int{
		{0, 0, 0, 0}, {0, 1, 2, 0}, {0, 2, 0, 0},
	}))
	println(maximumMinutes([][]int{
		{0, 0, 0, 0},
	}))
}
