package leetcode

import (
	"fmt"
	"testing"
)

func Test_findLadders(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	for _, v := range findLadders(beginWord, endWord, wordList) {

		fmt.Println(v)
	}
}
