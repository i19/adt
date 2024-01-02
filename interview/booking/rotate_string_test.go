package booking

import (
	"fmt"
	"testing"
)

func Test_rotateString(t *testing.T) {
	s := "abcde"
	for n := 0; n <= len(s); n++ {
		fmt.Printf("%d -> %s\n", n, rotateString(s, n))
	}
}
