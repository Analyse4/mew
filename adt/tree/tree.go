package tree

type Node struct {
	Value    interface{}
	Children []*Node
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return new(Tree)
}

// O(n*(1-(1/2)^h)/(1/2)) ?
func (t *Tree) Height() int {
	h := 0
	for n := t.root; n != nil; {
		h++
		if Size(n.Children[0]) >= Size(n.Children[1]) {
			n = n.Children[0]
		} else {
			n = n.Children[1]
		}
	}
	return h
}

func (t *Tree) ChangeRoot(r *Node) {
	t.root = r
}

// O(N)
func Size(n *Node) int {
	if n == nil {
		return 0
	}
	return Size(n.Children[0]) + 1 + Size(n.Children[1])
}

// TODO: generic
func (t *Tree) Insert(v interface{}) {
	if t.root == nil {
		t.root = &Node{v, make([]*Node, 2)}
		return
	}
	n := t.root
	for {
		if n.Value.(int) == v.(int) {
			return
		} else if n.Value.(int) > v.(int) {
			if n.Children[0] == nil {
				n.Children[0] = &Node{v, make([]*Node, 2)}
				return
			}
			n = n.Children[0]
		} else {
			if n.Children[1] == nil {
				n.Children[1] = &Node{v, make([]*Node, 2)}
				return
			}
			n = n.Children[1]
		}
	}
}
