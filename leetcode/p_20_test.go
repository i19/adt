package leetcode

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {
	inputs := []string{
		"()",
		")()",
		"([])",
		"([})",
	}
	for _, input := range inputs {
		fmt.Println(input, isValid(input))
	}
}
