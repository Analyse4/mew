package linkedlist

import "sync"

type Node struct {
	Key   interface{}
	Value interface{}
	Next  *Node
}

type Linkedlist struct {
	First *Node
	len   int
	mu    *sync.Mutex
}

func New() *Linkedlist {
	return &Linkedlist{nil, 0, new(sync.Mutex)}
}

func (q *Linkedlist) Insert(k, v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.IsEmpty() {
		q.First = &Node{k, v, nil}
		q.len++
		return
	}
	inode := q.First
	for ; inode != nil; inode = inode.Next {
		if inode.Key == k {
			inode.Value = v
		}
	}
	if inode == nil {
		inode = &Node{k, v, q.First}
		q.First = inode
		q.len++
	}
}

func (q *Linkedlist) Get(k interface{}) interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	inode := q.First
	for ; inode != nil; inode = inode.Next {
		if inode.Key == k {
			return inode.Value
		}
	}
	return nil
}

func (q *Linkedlist) Delete(k interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	inode := q.First
	for ; inode != nil; inode = inode.Next {
		if inode.Next.Key == k {
			nextnode := inode.Next.Next
			nextnode.Next = nextnode
			q.len--
		}
	}
}

func (q *Linkedlist) IsEmpty() bool {
	return q.First == nil
}

func (q *Linkedlist) Size() int {
	return q.len
}
