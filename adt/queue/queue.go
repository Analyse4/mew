package queue

type node struct {
	Key   interface{}
	Value interface{}
	next  *node
}

type queue struct {
	first *node
	end   *node
	len   int
}

func New() *queue {
	return new(queue)
}

func (q *queue) Enqueue(k, v interface{}) {
	n := &node{k, v, nil}
	if q.end == nil {
		q.first = n
		q.end = q.first
	}
	q.end.next = n
	q.end = n
	q.len++
}

func (q *queue) Dequeue() *node {
	n := q.first
	q.first = q.first.next
	q.len--
	return n
}

func (q *queue) IsEmpty() bool {
	if q.first == nil {
		return true
	}
	return false
}

func (q *queue) size() int {
	return q.len
}
