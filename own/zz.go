package own

import (
	"fmt"
	"sort"
)

func toB(n int) {
	fmt.Printf("%04b\n", n)
	sort.Slice(nil, func(i, j int) bool {})
}
