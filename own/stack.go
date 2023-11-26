package own

type stack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}

type arrayStack struct {
	data []interface{}
}

func newArrayStack() *arrayStack {
	return &arrayStack{data: make([]interface{}, 0, 100)}
}

func (s *arrayStack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *arrayStack) Pop() interface{} {
	v := s.Peek()
	s.data = s.data[0 : s.Size()-1]
	return v
}

func (s *arrayStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.Size()-1]
}

func (s *arrayStack) Size() int {
	return len(s.data)
}

func (s *arrayStack) IsEmpty() bool {
	return s.Size() == 0
}

type linkListStack struct {
	head *linkListNode
	size int
}

func newLinkListStack() *linkListStack {
	return &linkListStack{
		head: &linkListNode{
			next: nil,
		},
		size: 0,
	}
}
func (s *linkListStack) Push(v interface{}) {
	s.head.next = &linkListNode{
		value: v,
		next:  s.head.next,
	}
	s.size++
}

func (s *linkListStack) Pop() interface{} {
	v := s.Peek()
	if v != nil {
		s.head.next = s.head.next.next
		s.size--
	}
	return v
}

func (s *linkListStack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.head.next.value
}

func (s *linkListStack) Size() int {
	return s.size
}

func (s *linkListStack) IsEmpty() bool {
	return s.size == 0
}
