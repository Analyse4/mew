package linkedlist

import "sync"

type node struct {
	Key   interface{}
	Value interface{}
	next  *node
}

type Linkedlist struct {
	first *node
	len   int
	mu    *sync.Mutex
}

func New() *Linkedlist {
	return new(Linkedlist)
}

func (q *Linkedlist) Insert(k, v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	inode := q.first
	for ; inode.next != nil; inode = inode.next {
		if inode.Key == k {
			inode.Value = v
		}
	}
	if inode == nil {
		inode = &node{k, v, q.first}
		q.first = inode
		q.len++
	}
}

func (q *Linkedlist) Get(k interface{}) interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	inode := q.first
	for ; inode.next != nil; inode = inode.next {
		if inode.Key == k {
			return inode.Value
		}
	}
	return nil
}

func (q *Linkedlist) Delete(k interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	inode := q.first
	for ; inode.next != nil; inode = inode.next {
		if inode.next.Key == k {
			nextnode := inode.next.next
			nextnode.next = nextnode
			q.len--
		}
	}
}

func (q *Linkedlist) IsEmpty() bool {
	return q.first == nil
}

func (q *Linkedlist) Size() int {
	return q.len
}
