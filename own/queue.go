package own

type queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
	Size() int
	IsEmpty() bool
	IsFull() bool
}

type arrayQueue struct {
	data       []interface{}
	size       int
	cap        int
	head, tail int
}

func (q *arrayQueue) Enqueue(v interface{}) {
	if q.IsFull() {
		panic("full")
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
	q.size++
}

func (q *arrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size++
	return v
}

func (q *arrayQueue) Size() int {
	return q.size
}

func (q *arrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *arrayQueue) IsFull() bool {
	return q.size == q.cap
}

func newArrayQueue(size int) *arrayQueue {
	return &arrayQueue{data: make([]interface{}, size), cap: size}
}

type linkListQueue struct {
	head, tail *linkListNode
	size       int
	cap        int
}

func newLinkListQueue(size int) *linkListQueue {
	return &linkListQueue{cap: size}
}
func (q *linkListQueue) Enqueue(v interface{}) {
	if q.IsFull() {
		panic("full")
	}

	newNode := &linkListNode{value: v}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.head.next = newNode
		q.head = newNode
	}

	q.size++
}

func (q *linkListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	v := q.tail.value
	q.tail = q.tail.next
	q.size--
	return v
}

func (q *linkListQueue) Size() int {
	return q.size
}

func (q *linkListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *linkListQueue) IsFull() bool {
	return q.size == q.cap
}
