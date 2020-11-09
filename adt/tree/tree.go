package tree

type Tree struct {
	Value    interface{}
	Children []*Tree
}

// type Tree struct {
// 	root *Node
// }

func New() *Tree {
	return &Tree{nil, make([]*Tree, 2)}
}

// O(n*(1-(1/2)^h)/(1/2)) ?
// func (t *Tree) Height() int {
// 	h := 0
// 	for n := t.root; n != nil; {
// 		h++
// 		if Size(n.Children[0]) >= Size(n.Children[1]) {
// 			n = n.Children[0]
// 		} else {
// 			n = n.Children[1]
// 		}
// 	}
// 	return h
// }

// O(N)
func (t *Tree) Height() int {
	if t == nil {
		return 0
	}
	hl := 0
	hr := 0
	if t.Children[0] == nil {
		hl = 0
	} else {
		hl = t.Children[0].Height()
	}
	if t.Children[1] == nil {
		hr = 0
	} else {
		hr = t.Children[1].Height()
	}
	if hl >= hr {
		return hl + 1
	}
	return hr + 1
}

// O(N)
func Size(n *Tree) int {
	if n == nil {
		return 0
	}
	return Size(n.Children[0]) + 1 + Size(n.Children[1])
}

// TODO: generic
func (t *Tree) Insert(v interface{}) {
	if t.Value == nil {
		t.Value = v
		return
	}
	n := t
	for {
		if n.Value.(int) == v.(int) {
			return
		} else if n.Value.(int) > v.(int) {
			if n.Children[0] == nil {
				n.Children[0] = &Tree{v, make([]*Tree, 2)}
				return
			}
			n = n.Children[0]
		} else {
			if n.Children[1] == nil {
				n.Children[1] = &Tree{v, make([]*Tree, 2)}
				return
			}
			n = n.Children[1]
		}
	}
}

// O(N)
func GenerateMinimalHeightTree(l []int) *Tree {
	if len(l) == 0 {
		return nil
	}
	mid := len(l) / 2
	n := &Tree{l[mid], make([]*Tree, 2)}
	if len(l) == 1 {
		return n
	}
	n.Children[0] = GenerateMinimalHeightTree(l[:mid])
	if mid+1 > len(l)-1 {
		n.Children[1] = nil
		return n
	}
	n.Children[1] = GenerateMinimalHeightTree(l[mid+1:])

	return n
}
