package own

import (
	"fmt"
	"testing"
)

func Test_backTrackNQueen(t *testing.T) {
	for n := 4; n <= 5; n++ {
		res := backTrackNQueen(n)
		for _, v := range res {
			printCube(v, n)
			fmt.Println("--------------------")
		}
	}
}
