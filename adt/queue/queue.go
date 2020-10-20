package queue

type node struct {
	Key   interface{}
	Value interface{}
	next  *node
}

type Queue struct {
	first *node
	last  *node
	len   int
}

func New() *Queue {
	return new(Queue)
}

func (q *Queue) Enqueue(k, v interface{}) {
	oldlast := q.last
	q.last = &node{k, v, nil}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
	q.len++
}

func (q *Queue) Dequeue() *node {
	n := q.first
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.len--
	return n
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) Search(key interface{}) *node {
	return nil
}
