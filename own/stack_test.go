package own

import (
	"testing"
)

func Test_arrayStack(t *testing.T) {
	a := newLinkListStack()
	x := func(s stack) {
		println(s.IsEmpty())
		println(s.Size())
		s.Push(1)
		s.Push(2)
		println(s.IsEmpty())
		println(s.Size())
		println(s.Peek().(int))
		println(s.Pop().(int))
		println(s.Pop().(int))
		println(s.IsEmpty())
		println(s.Size())
	}
	x(a)
}
