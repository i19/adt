package own

import (
	"testing"
)

func Test_newArrayQueue(t *testing.T) {
	a := newArrayQueue(10)
	b := newLinkListQueue(10)
	x := func(s queue) {
		println(s.IsEmpty())
		println(s.Size())
		for i := 0; i < 10; i++ {
			s.Enqueue(i)
		}
		println(s.IsEmpty())
		println(s.Size())
		for i := 0; i < 10; i++ {
			println(s.Dequeue().(int))
		}
		println(s.IsEmpty())
		println(s.Size())
	}
	x(a)
	x(b)
}
