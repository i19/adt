package own

import (
	"testing"
)

func Test_newHashMapChaining(t *testing.T) {
	m := newHashMapChaining()
	m.put(0, "v1")
	m.put(4, "v2")
	m.extend()
	// for i := 0; i < 4; i++ {
	// 	m.put(i, fmt.Sprintf("v%d", i))
	// }
	m.print()
	// m.remove(2)
}
