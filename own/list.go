package own

type ListObject interface {
	Append(v interface{})
	Insert(offset int, v interface{})
	Delete(offset int)
	Update(offset int, v interface{})
	Get(offset int) interface{}
}

type sliceList struct {
	data []interface{}
}

func (s sliceList) Append(v interface{}) {
	s.data = append(s.data, v)
}

func (s sliceList) Insert(offset int, v interface{}) {
	if offset > len(s.data) {
		offset = len(s.data)
	}
	x := s.data[0:offset]
	x = append(x, v)
	x = append(x, s.data[offset:len(s.data)])
	s.data = x
}

func (s sliceList) Delete(offset int) {
	if offset >= len(s.data) {
		return
	}
	panic("implement me")
}

func (s sliceList) Update(offset int, v interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s sliceList) Get(offset int) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewListBySlice() *sliceList {
	return &sliceList{data: []interface{}{}}
}

type OneWayListNode struct {
	Value int
	Next  *OneWayListNode
}

func NewOneWayListNode(value int) *OneWayListNode {
	return &OneWayListNode{
		Value: value,
		Next:  nil,
	}
}

type TwoWayListNode struct {
	Value int
	Prev  *OneWayListNode
	Next  *OneWayListNode
}

func NewTwoWayListNode(value int) *TwoWayListNode {
	return &TwoWayListNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}
